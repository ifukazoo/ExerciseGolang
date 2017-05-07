package format

import (
	"bytes"
	"fmt"
	"testing"

	"../github"
)

func Test_SearchIssues(t *testing.T) {

	result := github.IssuesSearchResult{
		TotalCount: 1,
		Items: []*github.Item{
			&github.Item{
				Title:   "test title",
				Number:  17,
				HTMLURL: "http://test/17",
				State:   "closed",
				User: &github.User{
					Login:   "figaro",
					HTMLURL: "http://test/user",
				},
				MileStone: &github.MileStone{
					Title:   "alpha",
					HTMLURL: "http://test/milestone",
				},
			},
		},
	}
	buffer := bytes.NewBufferString("")
	OutputHTML(buffer, &result)
	fmt.Printf(buffer.String())
}
