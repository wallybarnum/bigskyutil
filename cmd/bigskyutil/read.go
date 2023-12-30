package bigskyutil

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/wallybarnum/bigskyutil/pkg/bigskyutil"

	"gitlab.com/gomidi/midi/v2"
	_ "gitlab.com/gomidi/midi/v2/drivers/rtmididrv"
)

var readCmd = &cobra.Command{
	Use:     "read",
	Aliases: []string{"rd"},
	Short:   "read file",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		read(args[0])
	},
}

func init() {
	rootCmd.AddCommand(readCmd)
}

var sig bool = false
var eof bool = false

func read(directory string) {

	defer midi.CloseDriver()

	in, err := midi.FindInPort("BigSkyMX")
	if err != nil {
		fmt.Println("can't find BigSkyMX")
		return
	}
	// listens to the in port and calls eachMessage for every message.
	// any running status bytes are converted and only complete messages are passed to the eachMessage.
	stop, err := midi.ListenTo(in, bigskyutil.EachMessage, midi.UseSysEx())
	if err != nil {
		fmt.Println("can't listen to in port")
		return
	}

	var out, _ = midi.FindOutPort("BigSkyMX")
	if err != nil {
		fmt.Println("can't find BigSkyMX")
		return
	}

	// creates a sender function to the out port
	send, err := midi.SendTo(out)
	if err != nil {
		fmt.Println("can't make send function")
		return
	}
	// make a dir request sysex message
	m, err := bigskyutil.ReadFileRequest(directory)
	if err != nil {
		fmt.Println("can't build dir request")
		return
	}

	err = send(m) // send the read request sysex message
	if err != nil {
		fmt.Println("can't send sysex message err = ", err)
		return
	}
	// wait for a response
	for bigskyutil.Eof == false {
		if bigskyutil.Sig == false {
			//fmt.Println("waiting for response")
			time.Sleep(time.Millisecond * 100)
		} else {
			bigskyutil.Sig = false
			// make a data block request sysex message
			m, err = bigskyutil.DataBlockRequest(256)
			err = send(m) // send the datablock sysex message
			if err != nil {
				fmt.Println("can't send data block requeste sysex message err = ", err)
				return
			}
		}
	}
	// stops listening
	stop()
}
/*
var eachMessage = func(msg midi.Message, timestampms int32) {
	if msg.Is(midi.RealTimeMsg) {
		// ignore realtime messages
		return
	}
	var channel, key, velocity, cc, val, prog uint8
	var bt []byte

	fmt.Println("RX msg")
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
		fmt.Println("RX sysex")
		sig = true

	default:
		fmt.Printf("%s\n", msg)
	}
}
*/

// example of preset directory listing
//{"dir" : [{"fn":"pst0.json","attr" : 0 },{"fn":"pst1.json","attr" : 0 },{"fn":"pst10.json","attr" : 0 },{"fn":"pst100.json","attr" : 0 },{"fn":"pst101.json","attr" : 0 },{"fn":"pst102.json","attr" : 0 },{"fn":"pst103.json","attr" : 0 },{"fn":"pst104.json","attr" : 0 },{"fn":"pst105.json","attr" : 0 },{"fn":"pst106.json","attr" : 0 },{"fn":"pst107.json","attr" : 0 },{"fn":"pst108.json","attr" : 0 },{"fn":"pst109.json","attr" : 0 },{"fn":"pst11.json","attr" : 0 },{"fn":"pst110.json","attr" : 0 },{"fn":"pst111.json","attr" : 0 },{"fn":"pst112.json","attr" : 0 },{"fn":"pst113.json","attr" : 0 },{"fn":"pst114.json","attr" : 0 },{"fn":"pst115.json","attr" : 0 },{"fn":"pst116.json","attr" : 0 },{"fn":"pst117.json","attr" : 0 },{"fn":"pst118.json","attr" : 0 },{"fn":"pst119.json","attr" : 0 },{"fn":"pst12.json","attr" : 0 },{"fn":"pst120.json","attr" : 0 },{"fn":"pst121.json","attr" : 0 },{"fn":"pst122.json","attr" : 0 },{"fn":"pst123.json","attr" : 0 },{"fn":"pst124.json","attr" : 0 } ] }
/*
  0  00 01 55 18 01 50 45 7B  22 64 69 72 22 20 3A 20  ..U..PE{"dir" : 
  16  5B 7B 22 66 6E 22 3A 22  70 73 74 30 2E 6A 73 6F  [{"fn":"pst0.jso
  32  6E 22 2C 22 61 74 74 72  22 20 3A 20 30 20 7D 2C  n","attr" : 0 },
  48  7B 22 66 6E 22 3A 22 70  73 74 31 2E 6A 73 6F 6E  {"fn":"pst1.json
  64  22 2C 22 61 74 74 72 22  20 3A 20 30 20 7D 2C 7B  ","attr" : 0 },{
  80  22 66 6E 22 3A 22 70 73  74 31 30 2E 6A 73 6F 6E  "fn":"pst10.json
  96  22 2C 22 61 74 74 72 22  20 3A 20 30 20 7D 2C 7B  ","attr" : 0 },{
 112  22 66 6E 22 3A 22 70 73  74 31 30 30 2E 6A 73 6F  "fn":"pst100.jso
 128  6E 22 2C 22 61 74 74 72  22 20 3A 20 30 20 7D 2C  n","attr" : 0 },
 144  7B 22 66 6E 22 3A 22 70  73 74 31 30 31 2E 6A 73  {"fn":"pst101.js
 160  6F 6E 22 2C 22 61 74 74  72 22 20 3A 20 30 20 7D  on","attr" : 0 }
 176  2C 7B 22 66 6E 22 3A 22  70 73 74 31 30 32 2E 6A  ,{"fn":"pst102.j
 192  73 6F 6E 22 2C 22 61 74  74 72 22 20 3A 20 30 20  son","attr" : 0 
 208  7D 2C 7B 22 66 6E 22 3A  22 70 73 74 31 30 33 2E  },{"fn":"pst103.
 224  6A 73 6F 6E 22 2C 22 61  74 74 72 22 20 3A 20 30  json","attr" : 0
 240  20 7D 2C 7B 22 66 6E 22  3A 22 70 73 74 31 30 34   },{"fn":"pst104
 256  2E 6A 73 6F 6E 22 2C 22  61 74 74 72 22 20 3A 20  .json","attr" : 
 272  30 20 7D 2C 7B 22 66 6E  22 3A 22 70 73 74 31 30  0 },{"fn":"pst10
 288  35 2E 6A 73 6F 6E 22 2C  22 61 74 74 72 22 20 3A  5.json","attr" :
 304  20 30 20 7D 2C 7B 22 66  6E 22 3A 22 70 73 74 31   0 },{"fn":"pst1
 320  30 36 2E 6A 73 6F 6E 22  2C 22 61 74 74 72 22 20  06.json","attr" 
 336  3A 20 30 20 7D 2C 7B 22  66 6E 22 3A 22 70 73 74  : 0 },{"fn":"pst
 352  31 30 37 2E 6A 73 6F 6E  22 2C 22 61 74 74 72 22  107.json","attr"
 368  20 3A 20 30 20 7D 2C 7B  22 66 6E 22 3A 22 70 73   : 0 },{"fn":"ps
 384  74 31 30 38 2E 6A 73 6F  6E 22 2C 22 61 74 74 72  t108.json","attr
 400  22 20 3A 20 30 20 7D 2C  7B 22 66 6E 22 3A 22 70  " : 0 },{"fn":"p
 416  73 74 31 30 39 2E 6A 73  6F 6E 22 2C 22 61 74 74  st109.json","att
 432  72 22 20 3A 20 30 20 7D  2C 7B 22 66 6E 22 3A 22  r" : 0 },{"fn":"
 448  70 73 74 31 31 2E 6A 73  6F 6E 22 2C 22 61 74 74  pst11.json","att
 464  72 22 20 3A 20 30 20 7D  2C 7B 22 66 6E 22 3A 22  r" : 0 },{"fn":"
 480  70 73 74 31 31 30 2E 6A  73 6F 6E 22 2C 22 61 74  pst110.json","at
 496  74 72 22 20 3A 20 30 20  7D 2C 7B 22 66 6E 22 3A  tr" : 0 },{"fn":
 512  22 70 73 74 31 31 31 2E  6A 73 6F 6E 22 2C 22 61  "pst111.json","a
 528  74 74 72 22 20 3A 20 30  20 7D 2C 7B 22 66 6E 22  ttr" : 0 },{"fn"
 544  3A 22 70 73 74 31 31 32  2E 6A 73 6F 6E 22 2C 22  :"pst112.json","
 560  61 74 74 72 22 20 3A 20  30 20 7D 2C 7B 22 66 6E  attr" : 0 },{"fn
 576  22 3A 22 70 73 74 31 31  33 2E 6A 73 6F 6E 22 2C  ":"pst113.json",
 592  22 61 74 74 72 22 20 3A  20 30 20 7D 2C 7B 22 66  "attr" : 0 },{"f
 608  6E 22 3A 22 70 73 74 31  31 34 2E 6A 73 6F 6E 22  n":"pst114.json"
 624  2C 22 61 74 74 72 22 20  3A 20 30 20 7D 2C 7B 22  ,"attr" : 0 },{"
 640  66 6E 22 3A 22 70 73 74  31 31 35 2E 6A 73 6F 6E  fn":"pst115.json
 656  22 2C 22 61 74 74 72 22  20 3A 20 30 20 7D 2C 7B  ","attr" : 0 },{
 672  22 66 6E 22 3A 22 70 73  74 31 31 36 2E 6A 73 6F  "fn":"pst116.jso
 688  6E 22 2C 22 61 74 74 72  22 20 3A 20 30 20 7D 2C  n","attr" : 0 },
 704  7B 22 66 6E 22 3A 22 70  73 74 31 31 37 2E 6A 73  {"fn":"pst117.js
 720  6F 6E 22 2C 22 61 74 74  72 22 20 3A 20 30 20 7D  on","attr" : 0 }
 736  2C 7B 22 66 6E 22 3A 22  70 73 74 31 31 38 2E 6A  ,{"fn":"pst118.j
 752  73 6F 6E 22 2C 22 61 74  74 72 22 20 3A 20 30 20  son","attr" : 0 
 768  7D 2C 7B 22 66 6E 22 3A  22 70 73 74 31 31 39 2E  },{"fn":"pst119.
 784  6A 73 6F 6E 22 2C 22 61  74 74 72 22 20 3A 20 30  json","attr" : 0
 800  20 7D 2C 7B 22 66 6E 22  3A 22 70 73 74 31 32 2E   },{"fn":"pst12.
 816  6A 73 6F 6E 22 2C 22 61  74 74 72 22 20 3A 20 30  json","attr" : 0
 832  20 7D 2C 7B 22 66 6E 22  3A 22 70 73 74 31 32 30   },{"fn":"pst120
 848  2E 6A 73 6F 6E 22 2C 22  61 74 74 72 22 20 3A 20  .json","attr" : 
 864  30 20 7D 2C 7B 22 66 6E  22 3A 22 70 73 74 31 32  0 },{"fn":"pst12
 880  31 2E 6A 73 6F 6E 22 2C  22 61 74 74 72 22 20 3A  1.json","attr" :
 896  20 30 20 7D 2C 7B 22 66  6E 22 3A 22 70 73 74 31   0 },{"fn":"pst1
 912  32 32 2E 6A 73 6F 6E 22  2C 22 61 74 74 72 22 20  22.json","attr" 
 928  3A 20 30 20 7D 2C 7B 22  66 6E 22 3A 22 70 73 74  : 0 },{"fn":"pst
 944  31 32 33 2E 6A 73 6F 6E  22 2C 22 61 74 74 72 22  123.json","attr"
 960  20 3A 20 30 20 7D 2C 7B  22 66 6E 22 3A 22 70 73   : 0 },{"fn":"ps
 976  74 31 32 34 2E 6A 73 6F  6E 22 2C 22 61 74 74 72  t124.json","attr
 992  22 20 3A 20 30 20 7D 20  5D 20 7D                 " : 0 } ] }    

*/
