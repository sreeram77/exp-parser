package main

import (
	"fmt"

	"github.com/sreeram77/exp-parser/ioutil"
)

func main() {

	// Read file
	fil := ioutil.New()

	t, err := fil.Read()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(t)
	// Call parser

	// Write response in new file
}
