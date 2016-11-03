//go:generate ../genem -n Emitter -o emitter.go $GOFILE

package event

import (
	"time"
)

// Point represents an pair of x and y coordinates.
type Point struct {
	x float32
	y float32
}

type clickEvent struct {
	point     Point
	clickedAt time.Time
}
