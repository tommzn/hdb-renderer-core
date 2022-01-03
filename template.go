package core

import (
	"path"
	"strings"
	"text/template"

	config "github.com/tommzn/go-config"
)

// NewFileTemplate returns a new template for given file.
func NewFileTemplate(filename string) Template {

	name := path.Base(filename)
	tmpl := template.Must(template.New(name).Funcs(templateFunctions()).ParseFiles(filename))
	return &FileTemplate{
		tmpl: tmpl,
	}
}

// NewFileTemplateFromConfig will create a template for a file defined in passed config.
func NewFileTemplateFromConfig(conf config.Config, templateDirConfigKey, templateFileConfigKey string) Template {

	templateDir := conf.Get(templateDirConfigKey, config.AsStringPtr("Undefined"))
	if !strings.HasSuffix(*templateDir, "/") {
		*templateDir = *templateDir + "/"
	}
	templateFile := conf.Get(templateFileConfigKey, config.AsStringPtr("Undefined"))
	return NewFileTemplate(*templateDir + *templateFile)
}

// RenderWith will execute assigned template with passed data.
func (fileTemplate *FileTemplate) RenderWith(data interface{}) (string, error) {

	builder := new(strings.Builder)
	err := fileTemplate.tmpl.Execute(builder, data)
	return builder.String(), err
}

// TemplateFunctions will return "add" and "subtract" functions to be used in templates.
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
