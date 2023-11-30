package testutils

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
)

func ParseInputFile(path string) []string {
	inputFile, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	var inputData []string
	for scanner.Scan() {
		inputData = append(inputData, scanner.Text())
	}

	return inputData
}

func ParseTestFile(fileName string) []string {
	return ParseInputFile(filepath.Join("../", "test", fileName))
}
