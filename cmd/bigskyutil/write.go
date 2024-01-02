package bigskyutil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
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
		data, err := os.ReadFile(args[0])
		if err != nil {
			fmt.Println("can't read source file, err = ", err)
			return
		}

	// check if JSON is valid
	valid := json.Valid(data)
	if valid == false {
		log.Println("JSON is not valid")
	}
	// compact JSON to single line
	compact := &bytes.Buffer{}
	json.Compact(compact, data)
	data = compact.Bytes()

	if len(data) > 0 {
		err = bigskyutil.WriteFile(args[1], data)
		if err != nil {
			fmt.Println("can't write file, err = ", err)
		}
		fmt.Println("wrote ", len(data), " bytes to ", args[1])
	} else {
		fmt.Println("error: empty file")
	}
		return
	},
}

func init() {
	rootCmd.AddCommand(writeCmd)
}

