package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	r := bytes.NewReader([]byte(`package event
import (
	"net"
	"sync"
	b "bytes"
)
type fooEvent struct {
	arg0, arg1 string
	arg2       b.Buffer
	arg3       *sync.Mutex
	IgnoredArg int
}
type IgnoredEvent struct {
	conn net.Conn
}`))

	actual, err := parse("test.go", r)
	if err != nil {
		t.Error(err)
	}
	expected := &EventEmitterParams{
		Package: "event",
		Imports: []*Import{
			&Import{
				Name: "",
				Path: "\"net\"",
			},
			&Import{
				Name: "",
				Path: "\"sync\"",
			},
			&Import{
				Name: "b",
				Path: "\"bytes\"",
			},
		},
		Events: []*Event{
			&Event{
				Name: "fooEvent",
				Params: []*EventParam{
					&EventParam{
						Names: []string{"arg0", "arg1"},
						Type:  "string",
					},
					&EventParam{
						Names: []string{"arg2"},
						Type:  "b.Buffer",
					},
					&EventParam{
						Names: []string{"arg3"},
						Type:  "*sync.Mutex",
					},
				},
			},
		},
	}
	assert.Equal(t, expected, actual)
}
