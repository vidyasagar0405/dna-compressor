package main

import (
	dnaencoding "dna-compressor/internals/dna_encoding"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	if len(os.Args) < 2 {
		log.Fatal("Please provide a file path")
	}

	filePath := os.Args[1]
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	dna := strings.ReplaceAll(string(data), "\n", "")
	dna = strings.ReplaceAll(dna, "\r", "")
	dna = strings.ReplaceAll(dna, " ", "")

	fmt.Printf("Read %d DNA characters from file\n", len(dna))

	// Encode the DNA sequence
	encoded, err := dnaencoding.Encode(dna)
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
