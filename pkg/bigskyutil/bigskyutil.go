package bigskyutil

import (
	"encoding/json"
	"fmt"

	"gitlab.com/gomidi/midi/gm"
	"gitlab.com/gomidi/midi/v2"
)

const fid byte = 0x18
const pid byte = 0x01
const xfer_opcode byte = 0x50

var xfer_cmd_opcode = []byte{0x00}
var xfer_data_opcode = []byte{0x01}
var xfer_data_ack_opcode = []byte{0x45}
var xfer_data_nack_opcode = []byte{0x46}

// example of a command to get a directory listing
//F0 00 01 55 18 01 50 00 7B 22 63 6D 64 22 3A 20 22 64 69 72 22 2C 20 22 70 61 74 68 22 3A 20 22 2F 64 65 76 22 7D F7

// file transfer protocol header (FID and PID are the same for all commands)
var hdr = []byte{0x00, 0x01, 0x55, fid, pid, xfer_opcode}


// the header for a command is:
// (F0) 00 01 55 [FID] [PID] 50 00 [JSON PAYLOAD] (F7) (SysEx F0 and F7 are not included in the header and will be added later)
var cmd_hdr = append(hdr, xfer_cmd_opcode...)

// the header for a data block is:
// F0 00 01 55 [FID] [PID] 50 01 [SIZE HI] [SIZE LO] [DATA] F7
//
// SIZE HI, SIZE LO, and DATA are not always needed depending on the command.
var data_hdr = append(hdr, xfer_data_opcode...)

//    data block request
//    F0 00 01 55 [FID] [PID] 50 01 [SIZE HI] [SIZE LO] F7
//    where SIZE is the number of 8-bit bytes to read
//    
//    e.g. to request 256 bytes
//    F0 00 01 55 18 01 50 01 01 00 F7

func DataBlockRequest(size int) ([]byte, error) {

    bt := append(data_hdr, byte(size>>8))
    bt = append(bt, byte(size))
    // make a sysex message, which will prepend F0 and append F7 to byte slice
	m := midi.SysEx(bt)
	//fmt.Println("sysex msg to send: ", m)
    return m, nil
}

func DirRequest(directory string) ([]byte, error ){
    /*
	   To get directory listing you will first send a command with a json payload of:
	   {
	    "cmd" : "dir",
	    "path" : "/path/to/directory"


	   followed by data blocks until a NACK is returned
	   F0 00 01 55 [FID] [PID] 50 01 F7
	*/

	type Message struct {
		Cmd  string `json:"cmd"`
		Path string `json:"path"`
	}

	data := Message{"dir", directory}
	// make a json object
	json_obj, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return nil, err
	}
	// json_obj is a byte array
	//fmt.Printf("json_obj: %s\n", json_obj)
    
    bt := append(cmd_hdr, xfer_cmd_opcode...)
    bt = append(cmd_hdr, json_obj...)
	// make a sysex message, which will prepend F0 and append F7 to byte slice
	m := midi.SysEx(bt)
	//fmt.Println("sysex msg to send: ", m)
    return m, err
}

var Sig bool = false
var Eof bool = false
var Rxbytes = []byte("") // TODO: use channel instead


var EachMessage = func(msg midi.Message, timestampms int32) {
	if msg.Is(midi.RealTimeMsg) {
		// ignore realtime messages
		return
	}

	var channel, key, velocity, cc, val, prog uint8
	var bt []byte

	//fmt.Println("RX msg")
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
		//fmt.Println("RX sysex: % X", bt)
		//fmt.Println("RX sysex")
		//DumpByteSlice(bt)
		ack, data, err := parseSysex(bt)
		if err != nil {
			fmt.Printf("ERROR: %s\n", err)
			return
		}
		if ack[0] == xfer_data_ack_opcode[0] {
			//fmt.Println("ACK")
			//DumpByteSlice(data)
			//fmt.Println(string(data))
			Rxbytes = append(Rxbytes, data...)
		} else if ack[0] == xfer_data_nack_opcode[0] {
			//fmt.Println("NACK")
			Eof = true
		} else {
			fmt.Println("not an ack or nack")
		}
		Sig = true

	default:
		fmt.Printf("%s\n", msg)
	}
}

func parseSysex(bt []byte) ([]byte, []byte, error){
	//fmt.Println("RX sysex: % X", bt)
	//fmt.Println("RX sysex")
	//DumpByteSlice(bt)

	// TODO: check for valid header...

	// check for ack or nack
	if bt[6] == xfer_data_ack_opcode[0] {
		//fmt.Println("ACK ", len(bt), " bytes" )
		data := []byte{}
		for i := 7; i < len(bt); i++ {
			if (i+1)%8 != 0 { // every 8th byte is 0x00 and not part of the data - why?
				data = append(data, bt[i])
			}
		}
		//DumpByteSlice(data)
		return xfer_data_ack_opcode, data, nil
	}else if bt[6] == xfer_data_nack_opcode[0] {
		//fmt.Println("NACK")
		return xfer_data_nack_opcode, nil, nil
	}else{
		fmt.Println("not an ack or nack")
	}
	return nil, nil, nil
}