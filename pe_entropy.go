package main

import (
	"debug/pe"
	"fmt"
	"log"
	"math"
	"os"
)

func calculateEntropy(data []byte) (float64, int) {
	entropy := 0.0
	totalBytes := float64(len(data))

	frequency := make(map[byte]float64)
	for _, b := range data {
		frequency[b]++
	}

	for _, count := range frequency {
		probability := count / totalBytes
		entropy -= probability * math.Log2(probability)
	}

	return entropy, len(data)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run entropy_segments.go <PE_file_path>")
		os.Exit(1)
	}

	peFilePath := os.Args[1]

	file, err := os.Open(peFilePath)
	if err != nil {
		log.Fatalf("Error opening file: %v\n", err)
	}
	defer file.Close()

	peFile, err := pe.NewFile(file)
	if err != nil {
		log.Fatalf("Error parsing PE file: %v\n", err)
	}

	for _, section := range peFile.Sections {
		fmt.Printf("Section: %s\n", section.Name)
		sectionData, err := section.Data()
		if err != nil {
			log.Fatalf("Error reading section data: %v\n", err)
		}

		entropy, byteCount := calculateEntropy(sectionData)
		fmt.Printf("Entropy: %f\n", entropy)
		fmt.Printf("Byte Count: %d\n", byteCount)
	}
}
