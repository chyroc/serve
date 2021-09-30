package internal

import (
	"bytes"
	_ "embed"
	"html/template"
)

//go:embed template/404.html
var Html404Template string

//go:embed template/dir.html
var htmlDirTemplate string

//go:embed template/error.html
var htmlErrorTemplate string

func BuildTemplate(templateContent string, data interface{}) string {
	t, err := template.New("").Parse(templateContent)
	if err != nil {
		panic(err)
	}
	buf := new(bytes.Buffer)
	err = t.Execute(buf, data)
	if err != nil {
		panic(err)
	}
	return buf.String()
}
