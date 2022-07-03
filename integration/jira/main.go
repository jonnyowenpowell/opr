package jira

import (
	"context"
	"fmt"

	"github.com/andygrunwald/go-jira"
	"github.com/pkg/errors"
)

type JiraUser struct {
	client *jira.Client
}

func User(baseUrl, username, token string) (*JiraUser, error) {
	tp := jira.BasicAuthTransport{
		Username: username,
		Password: token,
	}

	c, err := jira.NewClient(tp.Client(), baseUrl)
	if err != nil {
		return nil, err
	}

	return &JiraUser{
		client: c,
	}, nil
}

func IssueName(issueId string) string {
	return fmt.Sprintf("%s: An ticket", issueId)
}

func (j *JiraUser) GetIssueSummary(ctx context.Context, issueKey string) (string, error) {
	issue, _, err := j.client.Issue.GetWithContext(ctx, issueKey, &jira.GetQueryOptions{})
	if err != nil {
		return "", errors.Errorf("failed to get JIRA issue: %w", err)
	}

	return issue.Fields.Summary, nil
}

func (j *JiraUser) GetIssueLink(issueKey string) string {
	link := j.client.GetBaseURL()
	link.Path = fmt.Sprintf("browse/%s", issueKey)
	return link.String()
}
