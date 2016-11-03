package event

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestEmitter(t *testing.T) {
	e := NewEmitter()

	clickEventCh := make(chan Point, 10)
	listener := func(point Point, clickedAt time.Time) {
		clickEventCh <- point
	}

	e.AddClickEventListener(listener)
	e.EmitClickEvent(Point{x: 1, y: 2}, time.Now())
	assert.Equal(t, Point{x: 1, y: 2}, <-clickEventCh)
	e.EmitClickEvent(Point{x: 1, y: 2}, time.Now())
	assert.Equal(t, Point{x: 1, y: 2}, <-clickEventCh)

	e.RemoveClickEventListener(listener)
	e.EmitClickEvent(Point{x: 1, y: 2}, time.Now())
	assert.Equal(t, 0, len(clickEventCh))

	e.AddClickEventListenerOnce(listener)
	e.EmitClickEvent(Point{x: 1, y: 2}, time.Now())
	assert.Equal(t, Point{x: 1, y: 2}, <-clickEventCh)
	e.EmitClickEvent(Point{x: 1, y: 2}, time.Now())
	assert.Equal(t, 0, len(clickEventCh))

	e.AddClickEventListenerOnce(listener)
	e.RemoveClickEventListener(listener)
	e.EmitClickEvent(Point{x: 1, y: 2}, time.Now())
	assert.Equal(t, 0, len(clickEventCh))
}
