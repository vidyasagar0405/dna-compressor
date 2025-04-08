package utils

import "bytes"

func CleanInput(input []byte) []byte {
	cleaned := bytes.ReplaceAll(input, []byte("\n"), []byte(""))
	cleaned = bytes.ReplaceAll(cleaned, []byte("\r"), []byte(""))
	cleaned = bytes.ReplaceAll(cleaned, []byte(" "), []byte(""))
	return cleaned
}
