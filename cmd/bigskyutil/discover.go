package bigskyutil

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wallybarnum/bigskyutil/pkg/bigskyutil"
)

var discoverCmd = &cobra.Command{
    Use:   "discover",
    Aliases: []string{"disco"},
    Short:  "discover midi devices",
    Run: func(cmd *cobra.Command, args []string) {
		inports, outports := bigskyutil.Discover()
		fmt.Printf("In Ports:\n%s\n", inports)
		fmt.Printf("Out Ports:\n%s\n", outports)
    },
}

func init() {
    rootCmd.AddCommand(discoverCmd)
}

