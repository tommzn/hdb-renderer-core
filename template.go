package core

import (
	"errors"
	"strings"
	"text/template"
)

// NewFileTemplate returns a new template for given file.
func NewFileTemplate(filename string) Template {
	return &FileTemplate{
		tmpl: template.Must(template.ParseFiles(filename)),
	}
}

// NewEinkTemplate returns a new tempalte for given file to render content for ePaper displays.
func NewEinkTemplate(filename string, dataSource DataSource) UITemplate {
	return &EinkTemplate{
		Template:   NewFileTemplate(filename),
		dataSource: dataSource,
		anchor:     Point{X: 0, Y: 0},
	}
}

func (fileTemplate *FileTemplate) RenderWith(data interface{}) (string, error) {

	builder := new(strings.Builder)
	err := fileTemplate.tmpl.Execute(builder, data)
	return builder.String(), err
}

// SetAnchor for the top left corner of a canvas.
func (einkTemplate *EinkTemplate) SetAnchor(anchor Point) {
	einkTemplate.anchor = anchor
}

// GetSize returns height and width a rendered canvas uses.
func (einkTemplate *EinkTemplate) GetSize() Size {
	return Size{Height: 0, Width: 0}
}

// Render generated content based assigned position and event values.
func (einkTemplate *EinkTemplate) Render() (string, error) {

	event, err := einkTemplate.dataSource.Get()
	if err != nil {
		return "", err
	}
	if event == nil {
		return "", errors.New("Missing event data.")
	}
	data := einkTemplateData{
		Anchor: einkTemplate.anchor,
		Event:  event,
	}
	return einkTemplate.Template.RenderWith(data)
}
