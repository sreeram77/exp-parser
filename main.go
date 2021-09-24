package main

import (
	"os"

	"github.com/sreeram77/exp-parser/ioutil"
	"github.com/sreeram77/exp-parser/parser"
)

func main() {

	inputFilePath := os.Getenv("INPUT_PATH")
	if inputFilePath == "" {
		inputFilePath = "./file/input.yml"
	}

	outputFilePath := os.Getenv("OUTPUT_PATH")
	if outputFilePath == "" {
		outputFilePath = "./file/ouput.yml"
	}

	// Initialize file reader
	fileIO := ioutil.New(inputFilePath, outputFilePath)

	// Read file
	t, err := fileIO.Read()
	if err != nil {
		panic(err)
	}

	// Initialize parser
	yamlParser := parser.New()

	// Call parser
	processed, err := yamlParser.Parse(t)
	if err != nil {
		panic(err)
	}

	// Write response in new file
	err = fileIO.Write(processed)
	if err != nil {
		panic(err)
	}
}
