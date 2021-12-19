package core

import (
	"text/template"
)

// FileTemplate used to generate content based on file templates.
type FileTemplate struct {

	// tmpl, ref to Golang standard tempalate
	tmpl *template.Template
}

// EinkTemplate, can be used to render content for ePaper displays.
type EinkTemplate struct {

	// Template is used as source for templates.
	Template

	// dataSource used to get template data.
	dataSource DataSource

	// anchor is the upper left corner of rendered canvas.
	anchor Point
}

// Point is used to define a position based on y/> coordinates.
type Point struct {
	X, Y int64
}

// Size a canvas uses in a template.
type Size struct {
	Height, Width int64
}

type einkTemplateData struct {
	Anchor Point
	Event  interface{}
}
