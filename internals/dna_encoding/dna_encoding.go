package dnaencoding

import (
	"fmt"
	"strings"
)

const allowed = "atgcATGC"

var Encoding = map[rune]byte{
	'A': 0b00, // 00
	'T': 0b01, // 01
	'G': 0b10, // 10
	'C': 0b11, // 11
	'a': 0b00, // 00
	't': 0b01, // 01
	'g': 0b10, // 10
	'c': 0b11, // 11
}

// Each DNA base is represented with 2 bits
// A/a = 00, C/c = 01, G/g = 10, T/t = 11
func Encode(dna string) ([]byte, error) {
	byteCount := (len(dna) + 3) / 4

	encodedBytes := make([]byte, byteCount)

	for i, char := range dna {
		if !strings.ContainsRune(allowed, char) {
			return nil, fmt.Errorf("Not a DNA char: %v", string(char))
		}

		// Get the 2-bit value for this nucleotide
		value := Encoding[char]

		// Calculate which byte and position this nucleotide belongs to
		byteIndex := i / 4
		posInByte := i % 4

		// Shift the 2-bit value to the right position in the byte
		// positions: [6-7][4-5][2-3][0-1]
		shiftAmount := 6 - (posInByte * 2)
		value <<= shiftAmount

		// OR the value into the right position
		encodedBytes[byteIndex] |= value
	}

	return encodedBytes, nil
}



