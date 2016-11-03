package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func run(inputFileName, emitterTypeName, outputFileName string, print bool) {
	f, err := os.Open(inputFileName)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}

	parsed, err := parse(inputFileName, f)
	if err != nil {
		log.Fatal(err)
	}

	data, err := generate(emitterTypeName, inputFileName, parsed)
	if err != nil {
		log.Fatal(err)
	}

	if print {
		fmt.Print(string(data))
		return
	}

	if err := ioutil.WriteFile(outputFileName, data, 0644); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Generated", outputFileName)
}
