package utils

import (
	"bytes"
	"html/template"
	"github.com/qjw/git-notify/templates"
)

func ExcuteTemplate(tpl string, data interface{}) string {
	ttt := template.New(tpl)
	src := string(templates.MustAsset(tpl))
	ttt.Parse(src)
	var buf bytes.Buffer
	ttt.Execute(&buf, data)
	return buf.String()
}