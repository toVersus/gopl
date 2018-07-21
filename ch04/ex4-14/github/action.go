package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// SearchIssues queries the GitHub issue tracker.
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, ""))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	// We must close resp.Body on all execution paths.
	if resp.StatusCode != http.StatusOK {
		// defer is not used intentionally.
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

func get(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("can(t get %s: %s", url, resp.Status)
	}
	return resp, nil
}

func GetIssue(owner string, repo string, number string) (*Issue, error) {
	urlStr := strings.Join([]string{APIURL, "repos", owner, repo, "issues", number}, "/")
	resp, err := get(urlStr)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var issue Issue
	if err := json.NewDecoder(resp.Body).Decode(&issue); err != nil {
		return nil, err
	}
	return &issue, nil
}

func GetIssues(owner, repo string) ([]Issue, error) {
	urlStr := strings.Join([]string{APIURL, "repos", owner, repo, "issues"}, "/")
	resp, err := get(urlStr)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var issues []Issue
	if err := json.NewDecoder(resp.Body).Decode(&issues); err != nil {
		return nil, err
	}
	return issues, nil
}

func update(owner, repo, number string, fields map[string]string) (*Issue, error) {
	buf := &bytes.Buffer{}
	if err := json.NewEncoder(buf).Encode(fields); err != nil {
		return nil, err
	}

	client := &http.Client{}
	url := strings.Join([]string{APIURL, "repos", owner, repo, "issues", number}, "/")
	req, err := http.NewRequest("PATCH", url, buf)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to update issue: %s", resp.Status)
	}
	var issue Issue
	if err = json.NewDecoder(resp.Body).Decode(&issue); err != nil {
		return nil, err
	}
	return &issue, nil
}

func CloseIssue(owner, repo, number string) error {
	if _, err := update(owner, repo, number, map[string]string{"state": "closed"}); err != nil {
		return err
	}
	return nil
}

func OpenIssue(owner, repo, number string) error {
	if _, err := update(owner, repo, number, map[string]string{"state": "open"}); err != nil {
		return err
	}
	return nil
}

func UpdateIssue(owner, repo, number, title, body string) error {
	fields := map[string]string{
		"title": title,
		"body":  body,
	}
	if _, err := update(owner, repo, number, fields); err != nil {
		return err
	}
	return nil
}
