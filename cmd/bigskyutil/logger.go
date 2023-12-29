package bigskyutil

import (
	"fmt"

	"time"

	"github.com/spf13/cobra"

	"gitlab.com/gomidi/midi/v2"
	_ "gitlab.com/gomidi/midi/v2/drivers/rtmididrv"
)

var loggerCmd = &cobra.Command{
    Use:   "logger",
    Aliases: []string{"log"},
    Short:  "log rx midi",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
		log()
    },
}

func init() {
    rootCmd.AddCommand(loggerCmd)
}

func log() {
	defer midi.CloseDriver()
	
	in, err := midi.FindInPort("BigSkyMX")
	if err != nil {
		fmt.Println("can't find BigSkyMX")
		return
	}

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
}
