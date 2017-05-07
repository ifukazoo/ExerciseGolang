package github

import "testing"

func Test_SearchIssues(t *testing.T) {
	// https://github.com/golang/go/issues/227
	r, _ := SearchIssues([]string{"repo:golang/go", "is:closed", "support for daemonize", "227"})

	if r.TotalCount != 1 {
		t.Errorf("return[%v]", r.TotalCount)
	}
	if r.Items[0].User.Login != "gopherbot" {
		t.Errorf("return[%v]", r.Items[0].User.Login)
	}
	if r.Items[0].MileStone.Title != "Unplanned" {
		t.Errorf("return[%v]", r.Items[0].MileStone.Title)
	}
}
