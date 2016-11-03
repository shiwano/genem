# genem [![Build Status](https://secure.travis-ci.org/shiwano/genem.png?branch=master)](http://travis-ci.org/shiwano/genem)

> :sparkles: Generates event emitter code from Go types.

genem is a code generator. It finds event types named with `fooEvent` naming convention from the specified Go file, and generates event emitter code that contains methods with parameters defined by the event types.

## Installation

```bash
$ go get -u github.com/shiwano/genem
```

## Usage

```
genem

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
```

## Examples

Write a Go file that defined event types:

```go
//go:generate genem -n Emitter -o emitter.go $GOFILE

package event

type fooEvent struct {
	bar string
	baz int
}

type quxEvent struct {
	quux float32
}
```

And execute `go generate` command, then you will get the code like below:

```go
package event

type Emitter struct {}

func NewEmitter() *Emitter

func (_e *Emitter) EmitFooEvent(bar string, baz int)
func (_e *Emitter) AddFooEventListener(listener func(bar string, baz int))
func (_e *Emitter) AddFooEventListenerOnce(listener func(bar string, baz int))
func (_e *Emitter) RemoveFooEventListener(listener func(bar string, baz int))

func (_e *Emitter) EmitQuxEvent(quux float32)
func (_e *Emitter) AddQuxEventListener(listener func(quux float32))
func (_e *Emitter) AddQuxEventListenerOnce(listener func(quux float32))
func (_e *Emitter) RemoveQuxEventListener(listener func(quux float32))
```

## License

Copyright (c) 2016 Shogo Iwano
Licensed under the MIT license.
