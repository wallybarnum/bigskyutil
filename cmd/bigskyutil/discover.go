package bigskyutil

import (
	"fmt"

	"github.com/spf13/cobra"

	"gitlab.com/gomidi/midi/v2"
	_ "gitlab.com/gomidi/midi/v2/drivers/rtmididrv"
)

var discoverCmd = &cobra.Command{
    Use:   "discover",
    Aliases: []string{"disco"},
    Short:  "discover midi devices",
    Run: func(cmd *cobra.Command, args []string) {
		discover()
    },
}

func init() {
    rootCmd.AddCommand(discoverCmd)
}

func discover() {
	defer midi.CloseDriver()
	inports := midi.GetInPorts().String()
	fmt.Printf("In Ports:\n%s\n", inports)
	outports := midi.GetOutPorts().String()
	fmt.Printf("Out Ports:\n%s\n", outports)

}
