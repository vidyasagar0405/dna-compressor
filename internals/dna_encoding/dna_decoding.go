package dnaencoding

import (
	"fmt"
)

var decodingStatic = map[byte]string{
	27:  "atgc", // 00011011
	30:  "atcg", // 00011110
	39:  "agtc", // 00100111
	45:  "agct", // 00101101
	54:  "actg", // 00110110
	57:  "acgt", // 00111001
	75:  "tagc", // 01001011
	78:  "tacg", // 01001110
	99:  "tgac", // 01100011
	108: "tgca", // 01101100
	114: "tcag", // 01110010
	120: "tcga", // 01111000
	135: "gatc", // 10000111
	141: "gact", // 10001101
	147: "gtac", // 10010011
	156: "gtca", // 10011100
	177: "gcat", // 10110001
	180: "gcta", // 10110100
	198: "catg", // 11000110
	201: "cagt", // 11001001
	210: "ctag", // 11010010
	216: "ctga", // 11011000
	225: "cgat", // 11100001
	228: "cgta", // 11100100
}

func Decode(dna []byte) {

	for _, b := range dna {
		fmt.Printf("%08b: %s\n", b, decodingStatic[b])
	}
}
