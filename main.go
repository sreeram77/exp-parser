package main

import (
	"github.com/sreeram77/exp-parser/ioutil"
	"github.com/sreeram77/exp-parser/parser"
)

func main() {

	// Initialize file reader
	fileIO := ioutil.New()
	yamlParser := parser.New()

	// Read file
	t, err := fileIO.Read()
	if err != nil {
		panic(err)
	}

	// Call parser
	err = yamlParser.Parse(t)
	if err != nil {
		panic(err)
	}

	// Write response in new file
	err = fileIO.Write(t)
	if err != nil {
		panic(err)
	}
}
