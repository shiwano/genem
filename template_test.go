package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	params := &EventEmitterParams{
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
				Name: "",
				Path: "\"reflect\"",
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
	data, err := generate("Emitter", "emitter.go", params)
	if err != nil {
		t.Error(err)
	}
	actual := string(data)
	expected := `// Generated by: genem (https://github.com/shiwano/genem)

package event

import (
	"reflect"
	"sync"

	b "bytes"
)

// Emitter represents an event emitter.
type Emitter struct {
	fooEventMu            *sync.Mutex
	fooEventListeners     []func(arg0, arg1 string, arg2 b.Buffer, arg3 *sync.Mutex)
	fooEventListenersOnce []func(arg0, arg1 string, arg2 b.Buffer, arg3 *sync.Mutex)
}

// NewEmitter creates an event emitter.
func NewEmitter() *Emitter {
	return &Emitter{

		fooEventMu: new(sync.Mutex),
	}
}

// EmitFooEvent emits the specified event.
func (_e *Emitter) EmitFooEvent(arg0, arg1 string, arg2 b.Buffer, arg3 *sync.Mutex) {
	_e.fooEventMu.Lock()
	listeners := make([]func(arg0, arg1 string, arg2 b.Buffer, arg3 *sync.Mutex), len(_e.fooEventListeners))
	copy(listeners, _e.fooEventListeners)
	listenersOnce := _e.fooEventListenersOnce
	_e.fooEventListenersOnce = make([]func(arg0, arg1 string, arg2 b.Buffer, arg3 *sync.Mutex), 0)
	_e.fooEventMu.Unlock()
	for _, l := range listeners {
		l(arg0, arg1, arg2, arg3)
	}
	for _, l := range listenersOnce {
		l(arg0, arg1, arg2, arg3)
	}
}

// AddFooEventListener registers the specified event listener.
func (_e *Emitter) AddFooEventListener(listener func(arg0, arg1 string, arg2 b.Buffer, arg3 *sync.Mutex)) {
	_e.fooEventMu.Lock()
	_e.fooEventListeners = append(_e.fooEventListeners, listener)
	_e.fooEventMu.Unlock()
}

// AddFooEventListenerOnce registers the specified event listener that is invoked only once.
func (_e *Emitter) AddFooEventListenerOnce(listener func(arg0, arg1 string, arg2 b.Buffer, arg3 *sync.Mutex)) {
	_e.fooEventMu.Lock()
	_e.fooEventListenersOnce = append(_e.fooEventListenersOnce, listener)
	_e.fooEventMu.Unlock()
}

// RemoveFooEventListener removes the event listener previously registered.
func (_e *Emitter) RemoveFooEventListener(listener func(arg0, arg1 string, arg2 b.Buffer, arg3 *sync.Mutex)) {
	listenerPtr := reflect.ValueOf(listener).Pointer()
	_e.fooEventMu.Lock()
	listeners := _e.fooEventListeners[:0]
	for _, l := range _e.fooEventListeners {
		if reflect.ValueOf(l).Pointer() != listenerPtr {
			listeners = append(listeners, l)
		}
	}
	_e.fooEventListeners = listeners
	listenersOnce := _e.fooEventListenersOnce[:0]
	for _, l := range _e.fooEventListenersOnce {
		if reflect.ValueOf(l).Pointer() != listenerPtr {
			listenersOnce = append(listenersOnce, l)
		}
	}
	_e.fooEventListenersOnce = listenersOnce
	_e.fooEventMu.Unlock()
}
`
	assert.Equal(t, expected, actual)
}
