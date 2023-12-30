package bigskyutil

import (
	"github.com/spf13/cobra"
	"github.com/wallybarnum/bigskyutil/pkg/bigskyutil"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls", "dir"},
	Short:   "list directory contents",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		bigskyutil.List(args[0])
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}



