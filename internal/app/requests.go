package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type RepoSummary struct {
	FullName    string `json:"full_name"`
	Description string `json:"description"`
	Stars       int    `json:"stargazers_count"`
	OpenIssues  int    `json:"open_issues_count"`
	RepoURL     string `json:"html_url"`
	Language    string `json:"language"`
}

func getSinceDate(period string) string {
	var days int
	switch period {
	case "day":
		days = -1
	case "week":
		days = -7
	case "month":
		days = -30
	default:
		days = -7
	}
	return time.Now().AddDate(0, 0, days).Format("2006-01-02")
}

func FetchTrendingRepos(lang, period string, top int) ([]RepoSummary, error) {
	since := getSinceDate(period)
	query := fmt.Sprintf("language:%s created:>%s", lang, since)

	params := url.Values{}
	params.Add("q", query)
	params.Add("sort", "stars")
	params.Add("order", "desc")
	params.Add("per_page", fmt.Sprintf("%d", top))

	apiURL := "https://api.github.com/search/repositories?" + params.Encode()

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "github-explorer-cli")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GitHub API error: %s", resp.Status)
	}

	var result struct {
		Items []RepoSummary `json:"items"`
	}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&result); err != nil {
		return nil, err
	}

	return result.Items, nil
}
