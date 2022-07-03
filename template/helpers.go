package template

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/AlecAivazis/survey/v2"
	"github.com/jonnyowenpowell/opr/display"
)

func InputHelper(message string) (string, error) {
	p := &survey.Input{
		Message: message,
	}
	var s string

	err := survey.AskOne(p, &s)
	if err != nil {
		return "", err
	}

	return s, nil
}

func EditorHelper(message string) (string, error) {
	p := &survey.Editor{
		Message: message,
	}
	var s string

	err := survey.AskOne(p, &s)
	if err != nil {
		return "", err
	}

	return strings.TrimRight(s, "\r\n "), nil
}

var checklistTemplate string = `{{range $s, $check := .}}
- [{{if $check -}} x {{- else }} {{ end}}] {{ $s}} 
{{- end}}`

func ChecklistHelper(message string, items ...string) (string, error) {
	p := &survey.MultiSelect{
		Options: items,
		Message: message,
	}
	xs := []string{}

	err := survey.AskOne(p, &xs, display.Opts)
	if err != nil {
		return "", err
	}

	data := make(map[string]bool, len(items))
	for _, v := range items {
		data[v] = false
		for _, a := range xs {
			if a == v {
				data[v] = true
				break
			}
		}
	}

	t := template.Must(template.New("checklist").Parse(checklistTemplate))
	var b bytes.Buffer

	if err := t.Execute(&b, data); err != nil {
		return "", err
	}
	return b.String(), nil
}
