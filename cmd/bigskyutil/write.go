package bigskyutil

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/wallybarnum/bigskyutil/pkg/bigskyutil"
)

var writeCmd = &cobra.Command{
	Use:     "write <src file> <dst file>",
	Aliases: []string{"wr"},
	Short:   "write file to BigSkyMX",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		bytes, err := os.ReadFile(args[0])
		if err != nil {
			fmt.Println("can't read source file, err = ", err)
			return
		}
		// TODO: check if bytes is valid JSON

		if len(bytes) > 0 {
			err = bigskyutil.WriteFile(args[1], bytes)
			if err != nil {
				fmt.Println("can't write file, err = ", err)
			}
			fmt.Println("wrote ", len(bytes), " bytes to ", args[1])
		} else {
			fmt.Println("error: empty file")
		}
		return
	},
}

func init() {
	rootCmd.AddCommand(writeCmd)
}

