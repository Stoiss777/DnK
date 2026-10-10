/**
 * This code implements GraphQL query to retrieving the issue ID
 * and GraphQL mutation to delete issue by its ID.
 */
package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type queryResult struct {
    Data struct {
        Repository struct {
            Issue struct {
                ID string `json:"id"`
            } `json:"issue"`
        } `json:"repository"`
    } `json:"data"`
}


func graphQL(query, vars string) (*http.Response, error) {
	body := map[string]string {
		"query": query,
		"variables": vars,
	}
    data, err := json.Marshal(body)
    if err != nil {
        return nil, err
    }
    req, err := http.NewRequest(
        http.MethodPost,
        GraphQLURL,
        bytes.NewReader(data),
    )
	if err != nil {
	    return nil, err
	}

	client := &http.Client{}
    req.Header.Set("Accept", "application/vnd.github+json")
    req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer " + GetLastAccessToken())
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("GraphQL request failed: %s", resp.Status)
	}

	return resp, nil
}


func GetIssueID(params *IssueRequestParams) (string, error) {
    const query = `
		query($owner: String!, $repo: String!, $number: Int!) {
		  repository(owner: $owner, name: $repo) {
		    issue(number: $number) {
		      id
		    }
		  }
		}
	`
	variables, err := json.Marshal(params)
	if err != nil {
		return "", err
	}

    resp, err := graphQL(query, string(variables))
	if err != nil {
	    return "", err
	}

	var result *queryResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return "", err
	}

    return result.Data.Repository.Issue.ID, nil
}


func DeleteIssueByID(issueID string) error {
	const mutation = `
		mutation($issueId: ID!) {
		  deleteIssue(input: {
		    issueId: $issueId
		  }) {
		    repository {
		      name
		    }
		  }
		}
	`
    params := map[string]any{
        "issueId": issueID,
    }
	variables, err := json.Marshal(params)
	if err != nil {
		return err
	}

    _, err = graphQL(mutation, string(variables))
	return err
}
