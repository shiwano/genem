package main

import (
	"log"
	"path/filepath"

	"github.com/docopt/docopt-go"
)

const version = "0.1.0"

const usage = `genem

Usage:
  genem [options] <file>
  genem -h | --help
  genem --version

Options:
  -n, --emitter-name string      Specify the event emitter type name (default: EventEmitter).
  -o, --output-file-name string  Specify the output file name (default: event_emitter.go).
  -p, --print                    Print the generated code without file output.
  -h, --help                     Output help information.
  -v, --version                  Output version.
`

func main() {
	args, err := docopt.Parse(usage, nil, true, version, true)
	if err != nil {
		log.Fatal(err)
	}

	inputFileName, err := filepath.Abs(args["<file>"].(string))
	if err != nil {
		log.Fatal(err)
	}

	var emitterTypeName string
	if args["--emitter-name"] == nil {
		emitterTypeName = "EventEmitter"
	} else {
		emitterTypeName = args["--emitter-name"].(string)
	}

	var outputFileName string
	if args["--output-file-name"] == nil {
		outputFileName = filepath.Join(filepath.Dir(inputFileName), "event_emitter.go")
	} else {
		outputFileName = filepath.Join(filepath.Dir(inputFileName), args["--output-file-name"].(string))
	}

	run(inputFileName, emitterTypeName, outputFileName, args["--print"].(bool))
}
