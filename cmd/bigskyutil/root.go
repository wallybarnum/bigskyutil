package bigskyutil

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)
var version = "0.0.1"

var rootCmd = &cobra.Command{
    Use:  "bigskyutil",
	Version: version,
    Short: "bigskyutil - a simple CLI to transform and inspect strings",
    Long: `bigskyutil is a super fancy CLI (kidding)
   
One can use bigskyutil to modify or inspect strings straight from the terminal`,
    Run: func(cmd *cobra.Command, args []string) {

    },
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
        os.Exit(1)
    }
}