package wxpush

import (
	"bytes"
	"text/template"

	"github.com/superggfun/smoba/config"
)

func Push(markdown Markdown) error {
	tmpl, err := template.ParseFiles("static/template.md")
	if err != nil {
		panic(err)
	}
	buffer := new(bytes.Buffer)
	tmpl.Execute(buffer, markdown)
	err = pushplus(config.ReadFile().Pushplus, buffer.String())
	if err != nil {
		return err
	}
	return nil
}

func PushE(markdown Markdown) error {
	tmpl, err := template.ParseFiles("static/error.md")
	if err != nil {
		panic(err)
	}
	buffer := new(bytes.Buffer)
	tmpl.Execute(buffer, markdown)
	err = pushplus(config.ReadFile().Pushplus, buffer.String())
	if err != nil {
		return err
	}
	return nil
}
