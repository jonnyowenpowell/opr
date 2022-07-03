package template

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"html/template"
	"time"

	"github.com/jonnyowenpowell/opr/integration/jira"
)

func user() (*jira.JiraUser, error) {
	return jira.User(
		"https://mina-digital-limited.atlassian.net",
		"jonny@mina.co.uk",
		"")
}

func JiraIssueLinkHelper(issueKey string) (string, error) {
	u, err := user()
	if err != nil {
		return "", err
	}

	return u.GetIssueLink(issueKey), nil
}

func JiraIssueSummaryHelper(issueKey string) (string, error) {
	u, err := user()
	if err != nil {
		return "", err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	return u.GetIssueSummary(ctx, issueKey)
}

func Run() {
	fns := make(map[string]any)
	fns["askText"] = InputHelper
	fns["askMultiline"] = EditorHelper
	fns["askChecklist"] = ChecklistHelper

	fns["issueLink"] = JiraIssueLinkHelper
	fns["issueSummary"] = JiraIssueSummaryHelper

	tmpl := template.Must(template.New("test").Funcs(fns).Parse(`Title: MINA-1565: {{issueSummary "MINA-1565"}}

## Description

{{askMultiline "Description"}}

## Related Issues

- [MINA-1565]({{issueLink "MINA-1565"}})
- [MINA-1460]({{issueLink "MINA-1460"}})

## Pre-review Checklist
{{ askChecklist "Pre-review checklist"
"a"
"b"
"c"}}
`))

	var b bytes.Buffer
	err := tmpl.Execute(&b, "no data needed")
	if err != nil {
		fmt.Printf("Failed to render template: %s", errors.Unwrap(err))
		return
	}

	fmt.Println("******")
	fmt.Print(b.String())
	fmt.Println("******")
}
