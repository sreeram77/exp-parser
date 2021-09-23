package main

import (
	"github.com/sreeram77/exp-parser/ioutil"
	"github.com/sreeram77/exp-parser/parser"
)

func main() {

	// Read file
	fil := ioutil.New()
	yamlParser := parser.New()

	t, err := fil.Read()
	if err != nil {
		panic(err)
	}

	// Call parser
	err = yamlParser.Parse(t)
	if err != nil {
		panic(err)
	}

	// Write response in new file
}
