package dnaencoding

import (
	"fmt"
)

var Encoding = map[byte]byte{
	65:  0b00, // 00, A
	84:  0b01, // 01, T
	71:  0b10, // 10, G
	67:  0b11, // 11, C
	97:  0b00, // 00, a
	116: 0b01, // 01, t
	103: 0b10, // 10, g
	99:  0b11, // 11, c
}

// Each DNA base is represented with 2 bits
// A/a = 00, C/c = 01, G/g = 10, T/t = 11
func Encode(dna []byte) ([]byte, error) {
	byteCount := (len(dna) + 3) / 4

	encodedBytes := make([]byte, byteCount)

	for i, b := range dna {

		value, ok := Encoding[b]
		if !ok {
			return nil, fmt.Errorf("invalid DNA base: %s\n", string(b))
		}

		// Calculate which byte and position this nucleotide belongs to
		byteIndex := i / 4
		posInByte := i % 4

		// Shift the 2-bit value to the right position in the byte
		// positions: [7-6][5-4][3-2][1-0]
		shiftAmount := 6 - (posInByte * 2)
		value <<= shiftAmount

		// OR the value into the right position
		encodedBytes[byteIndex] |= value
	}

	return encodedBytes, nil
}
