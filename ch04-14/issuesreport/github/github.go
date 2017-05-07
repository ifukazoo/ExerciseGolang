package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// IssuesURL Issues 検索用 api
const IssuesURL = "https://api.github.com/search/issues"

// IssuesSearchResult 検索結果 全体
type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Item
}

// Item 検索結果 Itemフィールド
type Item struct {
	HTMLURL   string `json:"html_url"`
	Number    int
	Title     string
	State     string
	User      *User
	MileStone *MileStone
}

// User 検索結果 Userフィールド
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

// MileStone 検索結果 MileStoneフィールド
type MileStone struct {
	Title   string
	HTMLURL string `json:"html_url"`
}

// SearchIssues github.comへissuesの問い合わせ
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
