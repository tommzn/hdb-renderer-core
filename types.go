package core

import (
	"text/template"
	"time"
)

// FileTemplate used to generate content based on file templates.
type FileTemplate struct {

	// tmpl, ref to Golang standard tempalate
	tmpl *template.Template
}

// Point is used to define a position based on y/> coordinates.
type Point struct {
	X, Y int
}

// Size a canvas uses in a template.
type Size struct {
	Height, Width int
}

// Spacing defines top, left right and bottom distance to other elements.
type Spacing struct {
	Top, Right, Bottom, Left int
}

// EventTypeName is used to identify events. Could be type name of an event, given by
// refect.TypeOf().String(). Maybe it's ectended by a suffix to distiguish different variants of
// an event.
type eventTypeName string

// EventTimestampManager manages timestamp for latest events.
type EventTimestampManager struct {
	timestamps map[eventTypeName]time.Time
}
