package main

import (
	"fmt"

	"github.com/sreeram77/exp-parser/ioutil"
	"github.com/sreeram77/exp-parser/parser"
)

func main() {

	// Read file
	fil := ioutil.New()
	yamlParser := parser.New()

	t, err := fil.Read()
	if err != nil {
		fmt.Println(err)
	}

	// Call parser
	yamlParser.Parse(t)

	// Write response in new file
}
