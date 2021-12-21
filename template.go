package core

import (
	"strings"
	"text/template"
)

// NewFileTemplate returns a new template for given file.
func NewFileTemplate(filename string) Template {
	return &FileTemplate{
		tmpl: template.Must(template.ParseFiles(filename)),
	}
}

func (fileTemplate *FileTemplate) RenderWith(data interface{}) (string, error) {

	builder := new(strings.Builder)
	err := fileTemplate.tmpl.Execute(builder, data)
	return builder.String(), err
}
