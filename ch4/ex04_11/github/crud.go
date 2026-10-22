/**
 * CRUD functions
 * 
 * Github API Manual: https://docs.github.com/en/rest/issues/issues?apiVersion=2026-03-10
 * 
 */
package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)


func GetIssues(params *IssueRequestParams) ([]*Issue, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf(GetIssuesURL, params.Owner, params.Repo),
		nil,
	)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("Authorization", "Bearer " + GetLastAccessToken())
	req.Header.Add("X-GitHub-Api-Version", GitHubAPIVer)

	client := &http.Client{}
	resp, err := client.Do(req)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("Request failed: %s\n", resp.Status)
	}

	var issues []*Issue
	if err := json.NewDecoder(resp.Body).Decode(&issues); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return issues, nil
}


func GetIssue(params *IssueRequestParams) (*Issue, error) {
	req, err := http.NewRequest(
	    http.MethodGet,
	    fmt.Sprintf(GetIssueURL, params.Owner, params.Repo, params.Number),
	    nil,
	)
	if err != nil {
	    return nil, err
	}

	client := &http.Client{}
	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("Authorization", "Bearer " + GetLastAccessToken())
	req.Header.Add("X-GitHub-Api-Version", GitHubAPIVer)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("Request failed: %s", resp.Status)
	}

	var issue *Issue
	if err := json.NewDecoder(resp.Body).Decode(&issue); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return issue, nil
}


func CreateIssue(params *IssueRequestParams, data *IssueRequestData) (*Issue, error) {
	body, err := json.Marshal(data)
	if err != nil {
	    return nil, err
	}
	req, err := http.NewRequest(
	    http.MethodPost,
	    fmt.Sprintf(CreateIssueURL, params.Owner, params.Repo),
	    bytes.NewReader(body),
	)
	if err != nil {
	    return nil, err
	}

	client := &http.Client{}
	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("Authorization", "Bearer " + GetLastAccessToken())
	req.Header.Add("X-GitHub-Api-Version", GitHubAPIVer)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusCreated {
		resp.Body.Close()
		return nil, fmt.Errorf("Request failed: %s", resp.Status)
	}

	var issue *Issue
	if err := json.NewDecoder(resp.Body).Decode(&issue); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return issue, nil
}


func UpdateIssue(params *IssueRequestParams, data *IssueRequestData) (*Issue, error) {
	body, err := json.Marshal(data)
	if err != nil {
	    return nil, err
	}
	req, err := http.NewRequest(
	    http.MethodPost,
	    fmt.Sprintf(UpdateIssueURL, params.Owner, params.Repo, params.Number),
	    bytes.NewReader(body),
	)
	if err != nil {
	    return nil, err
	}

	client := &http.Client{}
	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("Authorization", "Bearer " + GetLastAccessToken())
	req.Header.Add("X-GitHub-Api-Version", GitHubAPIVer)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("Request failed: %s", resp.Status)
	}

	var issue *Issue
	if err := json.NewDecoder(resp.Body).Decode(&issue); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return issue, nil
}

/*
 * The current GitHub API does not have a function to delete an issue;
 * however, such a function exists in the GraphQL interface.
 *
 * @see graphql.go
 */
func DeleteIssue(params *IssueRequestParams) error {
	id, err := GetIssueID(params)
	if err != nil {
		return err
	}

	return DeleteIssueByID(id)
}