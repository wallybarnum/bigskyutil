package bigskyutil

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/wallybarnum/bigskyutil/pkg/bigskyutil"
)

var readCmd = &cobra.Command{
	Use:     "read <src file> <dst outfile>",
	Aliases: []string{"rd"},
	Short:   "read file from BigSkyMX",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		bytes, err := bigskyutil.ReadFile(args[0])
		if err != nil {
			fmt.Println("can't read file, err = ", err)
		}
		if len(bytes) > 0 {
			f, err := os.Create(args[1])
			check(err)
			defer f.Close()
			err = os.WriteFile(args[1], bytes, 0644)
			check(err)
			f.Sync()
			fmt.Println("wrote ", len(bytes), " bytes to ", args[1])
			fmt.Println("readback: ")
			dat, err := os.ReadFile(args[1])	
		    check(err)
    		fmt.Print(string(dat))
		} else {
			fmt.Println("no bytes returned")
		}
	},
}

func init() {
	rootCmd.AddCommand(readCmd)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
