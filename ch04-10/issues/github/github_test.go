package github

import (
	"fmt"
	"testing"
)

func Test_SearchIssues(t *testing.T) {
	v, err := SearchIssues([]string{"repo:golang/go", "is:open", "json", "decoder"})
	fmt.Println(v)
	fmt.Println(err)
}
