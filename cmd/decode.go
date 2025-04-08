/*
Copyright © 2025 Vidyasagar Gopi vidyasagar0405@gmail.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	dnaencoding "dna-compressor/internals/dna_encoding"
	"dna-compressor/internals/utils"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// decodeCmd represents the decode command
var decodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "Decodes the encoded output",
	Long:  `Decodes the encoded output`,
	Run: func(cmd *cobra.Command, args []string) {

		filePath := args[0]

		data, err := os.ReadFile(filePath)
		if err != nil {
			log.Fatalf("Error reading file: %v", err)
		}

		outVer, err := dnaencoding.Decode(utils.CleanInput(data))
		if err != nil {
			log.Fatal("Cant decode: ", err)
		}

		outputPath := filePath + ".decoded"
		err = os.WriteFile(outputPath, []byte(outVer), 0644)
		if err != nil {
			log.Fatalf("Error writing decoded data: %v", err)
		}

		fmt.Printf("Decoded data written to %s\n", outputPath)
	},
}

func init() {
	rootCmd.AddCommand(decodeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// decodeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// decodeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
