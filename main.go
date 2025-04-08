package main

import (
	dnaencoding "dna-compressor/internals/dna_encoding"
	"fmt"
	"log"
	"os"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Print(USAGE)
		os.Exit(1)
	}

	cmd := os.Args[1]

	filePath := os.Args[2]
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	if cmd == "encode" {
		// Encode the DNA sequence
		encoded, err := dnaencoding.Encode(cleanInput(data))
		if err != nil {
			log.Fatalf("Error encoding DNA: %v", err)
		}

		outputPath := filePath + ".encoded"
		err = os.WriteFile(outputPath, encoded, 0644)
		if err != nil {
			log.Fatalf("Error writing encoded data: %v", err)
		}

		fmt.Printf("Encoded data written to %s\n", outputPath)
	}

	if cmd == "decode" {
		outVer, err := dnaencoding.Decode(cleanInput(data))
		if err != nil {
			log.Fatal("Cant decode: ", err)
		}

		outputPath := filePath + ".decoded"
		err = os.WriteFile(outputPath, []byte(outVer), 0644)
		if err != nil {
			log.Fatalf("Error writing decoded data: %v", err)
		}

		fmt.Printf("Decoded data written to %s\n", outputPath)
	}

}
