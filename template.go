package core

import (
	"path"
	"strings"
	"text/template"
)

// NewFileTemplate returns a new template for given file.
func NewFileTemplate(filename string) Template {

	name := path.Base(filename)
	tmpl := template.Must(template.New(name).Funcs(templateFunctions()).ParseFiles(filename))
	return &FileTemplate{
		tmpl: tmpl,
	}
}

func (fileTemplate *FileTemplate) RenderWith(data interface{}) (string, error) {

	builder := new(strings.Builder)
	err := fileTemplate.tmpl.Execute(builder, data)
	return builder.String(), err
}

func templateFunctions() template.FuncMap {
	return template.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
		"subtract": func(a, b int) int {
			return a - b
		},
	}
}
