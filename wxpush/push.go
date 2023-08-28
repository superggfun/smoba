package wxpush

import (
	"bytes"
	"text/template"

	"github.com/superggfun/smoba/config"
)

// pushTemplate parses template from the provided filename, execute it with the provided markdown
// data and push the result using the pushplus function.
func pushTemplate(filename string, markdown Markdown) error {
	tmpl, err := template.ParseFiles(filename)
	if err != nil {
		return err
	}

	buffer := new(bytes.Buffer)
	if err := tmpl.Execute(buffer, markdown); err != nil {
		return err
	}

	cfg, err := config.ReadConfigFile("config.json") // code/config.json
	if err != nil {
		return err
	}

	if err := pushplus(cfg.Wxpush.Pushplus, buffer.String()); err != nil {
		return err
	}

	return nil
}

func Push(markdown Markdown) error {
	return pushTemplate("static/template.md", markdown) // replace with your actual template file path
}

func PushE(markdown Markdown) error {
	return pushTemplate("static/error.md", markdown) // replace with your actual error template file path
}
