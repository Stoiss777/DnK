package github

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/**
 *  https://docs.github.com/en/rest/issues/issues?apiVersion=2026-03-10
 */
func GetBugs(owner, name string) ([]Issue, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf(GetIssuesURL + "?labels=bug", owner, name),
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

	var issues []Issue
	if err := json.NewDecoder(resp.Body).Decode(&issues); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return issues, nil
}


/**
 *  https://docs.github.com/en/rest/issues/milestones?apiVersion=2026-03-10
 */
func GetMilestones(owner, name string) ([]Milestone, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf(GetMilestonesURL, owner, name),
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

	var milestones []Milestone
	if err := json.NewDecoder(resp.Body).Decode(&milestones); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return milestones, nil
}


func GetContributors(owner, name string) ([]User, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf(GetContributorsURL, owner, name),
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

	var contributors []User
	if err := json.NewDecoder(resp.Body).Decode(&contributors); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return contributors, nil
}