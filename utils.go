package main

import (
	"bytes"
	"fmt"
	"os"
)

var USAGE = fmt.Sprintf(`
    Usage: %s <cmd> <input_file>

    cmd:

    encode - Encodes 4 DNA nucleotides in to a single byte
    decode - Decodes the encoded output

`, os.Args[0])

func cleanInput(input []byte) []byte {
	cleaned := bytes.ReplaceAll(input, []byte("\n"), []byte(""))
	cleaned = bytes.ReplaceAll(cleaned, []byte("\r"), []byte(""))
	cleaned = bytes.ReplaceAll(cleaned, []byte(" "), []byte(""))
	return cleaned
}
