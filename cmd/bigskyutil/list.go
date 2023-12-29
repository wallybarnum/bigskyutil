package bigskyutil

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"gitlab.com/gomidi/midi/gm"
	"gitlab.com/gomidi/midi/v2"
	_ "gitlab.com/gomidi/midi/v2/drivers/rtmididrv"
)

const fid byte = 0x18
const pid byte = 0x01
const xfer_opcode byte = 0x50

var xfer_cmd_opcode = []byte{0x00}
var xfer_data_opcode = []byte{0x01}

// example of a command to get a directory listing
//F0 00 01 55 18 01 50 00 7B 22 63 6D 64 22 3A 20 22 64 69 72 22 2C 20 22 70 61 74 68 22 3A 20 22 2F 64 65 76 22 7D F7

// file transfer protocol header (FID and PID are the same for all commands)
var hdr = []byte{0x00, 0x01, 0x55, fid, pid, xfer_opcode}

// the header for a command is:
// F0 00 01 55 [FID] [PID] 50 00 [JSON PAYLOAD] F7 (sx F0 and F7 are not included in the header and will be added later)
var cmd_hdr = append(hdr, xfer_cmd_opcode...)

// the header for a data block is:
// F0 00 01 55 [FID] [PID] 50 01 [SIZE HI] [SIZE LO] [DATA] F7
//
// SIZE HI, SIZE LO, and DATA are not always needed depending on the command.
var data_hdr = append(hdr, xfer_data_opcode...)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "list directory contents",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		list(args[0])
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}


func list(directory string) {

	/*
	   To get directory listing you will first send a command with a json payload of:
	   {
	    "cmd" : "dir",
	    "path" : "/path/to/directory"


	   followed by data blocks until a NACK is returned
	   F0 00 01 55 [FID] [PID] 50 01 F7
	*/

	// build midi sysex message
	type Message struct {
		Cmd  string `json:"cmd"`
		Path string `json:"path"`
	}
	data := Message{"dir", directory}
	// make a json object
	json_obj, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return
	}
	// json_obj is a byte array
	fmt.Printf("json_obj: %s\n", json_obj)

	bt := append(cmd_hdr, json_obj...)
	// make a sysex message, which will append F0 and add F7 to the end of byte slice
	m := midi.SysEx(bt)
	fmt.Println("msg to send: ", m)

	//the header for a data block is:
	//F0 00 01 55 [FID] [PID] 50 01 [SIZE HI] [SIZE LO] [DATA] F7
	//
	// SIZE HI, SIZE LO, and DATA are not always needed depending on the command.
	//	fmt.Println(data_hdr)// string = hdr + xfer_data_opcode + " "
	/*
		//cmd = cmd_hdr + _json2str(json_data) + " F7"
		//cmd = cmd.replace(' ', '')
		//bytes_obj = bytes.fromhex(cmd)

		//with open('getDir.syx', 'wb') as file:
		//    file.write(bytes_obj)
		//return cmd

		//examples:
		//getDirCmd("/dev")
		//getDirCmd("/presets")
	*/
	defer midi.CloseDriver()

	// var in = InByName("my midi keyboard")
	in, err := midi.FindInPort("BigSkyMX")
	//in, err := midi.FindInPort("Logic Pro Virtual Out")
	//in, err := midi.FindInPort("MIDI Monitor (Untitled)")
	if err != nil {
		//fmt.Println("can't find MIDI Monitor In port")
		fmt.Println("can't find BigSkyMX")
		return
	}
	// listens to the in port and calls eachMessage for every message.
	// any running status bytes are converted and only complete messages are passed to the eachMessage.
	stop, err := midi.ListenTo(in, eachMessage, midi.UseSysEx())
	if err != nil {
		fmt.Println("can't listen to in port")
		return
	}

	var out, _ = midi.FindOutPort("BigSkyMX")
	//var out, _ = midi.FindOutPort("MIDI Monitor (Untitled)")
	// var out = OutByName("my synth")
	if err != nil {
		//fmt.Println("can't find Midi Monitor Out port")
		fmt.Println("can't find BigSkyMX")
		return
	}

	// creates a sender function to the out port
	send, err := midi.SendTo(out)
	if err != nil {
		fmt.Println("can't make send function")
		return
	}


	{ // send some messages
		//	send(NoteOn(0, Db(5), 100))
		//	send(NoteOff(0, Db(5)))
		//	send(Pitchbend(0, -12))
		//	send(ProgramChange(1, gm.Instr_AcousticBass.Value()))
		//	send(ControlChange(2, FootPedalMSB, On))
		fmt.Println("sending sysex message")
		err = send(m) // send the sysex message
		if err != nil {
			fmt.Println("can't send sysex message err = ", err)
			return
		}
	}
	time.Sleep(time.Second * 15)
	// stops listening
	stop()

	//func SendTo(outPort drivers.Out) (func(msg Message) error, error)
	/*
		stop, err := midi.ListenTo(in, func(msg midi.Message, timestampms int32) {
			var bt []byte
			var ch, key, vel uint8
			switch {
			case msg.GetSysEx(&bt):
				//fmt.Printf("RX sysex: % X\n", bt)
				fmt.Printf("RX sysex:\n")
				dumpByteSlice(bt)
			case msg.GetNoteStart(&ch, &key, &vel):
				fmt.Printf("RX starting note %s on channel %v with velocity %v\n", midi.Note(key), ch, vel)
			case msg.GetNoteEnd(&ch, &key):
				fmt.Printf("RX ending note %s on channel %v\n", midi.Note(key), ch)
			default:
				// ignore
			}
		}, midi.UseSysEx())

		if err != nil {
			fmt.Printf("ERROR: %s\n", err)
			return
		}

		time.Sleep(time.Second * 15)

		stop()
	*/
}

var eachMessage = func(msg midi.Message, timestampms int32) {
	if msg.Is(midi.RealTimeMsg) {
		// ignore realtime messages
		return
	}
	var channel, key, velocity, cc, val, prog uint8
	var bt []byte

	switch {

	// is better, than to use GetNoteOn (handles note on messages with velocity of 0 as expected)
	case msg.GetNoteStart(&channel, &key, &velocity):
		fmt.Printf("note started channel: %v key: %v (%s) velocity: %v\n", channel, key, midi.Note(key), velocity)

	// is better, than to use GetNoteOff (handles note on messages with velocity of 0 as expected)
	case msg.GetNoteEnd(&channel, &key):
		fmt.Printf("note ended channel: %v key: %v (%s)\n", channel, key, midi.Note(key))

	case msg.GetControlChange(&channel, &cc, &val):
		fmt.Printf("control change %v (%s) channel: %v value: %v\n", cc, midi.ControlChangeName[cc], channel, val)

	case msg.GetProgramChange(&channel, &prog):
		fmt.Printf("program change %v (%s) channel: %v\n", prog, gm.Instr(prog), channel)

	case msg.GetSysEx(&bt):
		fmt.Printf("sysex: % X\n", bt)

	default:
		fmt.Printf("%s\n", msg)
	}
}
