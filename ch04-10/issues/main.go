package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"./github"
)

func filterBy(result *github.IssuesSearchResult, f func(*github.Issue) bool) *[]*github.Issue {
	var filterd []*github.Issue
	for _, v := range result.Items {
		if f(v) {
			filterd = append(filterd, v)
		}
	}
	return &filterd
}
func showIssue(issues *[]*github.Issue, title string) {
	fmt.Println(title)
	for _, item := range *issues {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.CreatedAt)
	}
}

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)

	/*
	 *以上,未満の厳密な比較はしていない
	 */
	dayOfOneMonthAgo := time.Now().AddDate(0, -1, 0)
	dayOfOneYearsAgo := time.Now().AddDate(-1, 0, 0)

	var filterd *[]*github.Issue

	filterd = filterBy(result, func(i *github.Issue) bool {
		return i.CreatedAt.After(dayOfOneMonthAgo)
	})
	showIssue(filterd, "Report less than one month==================================")

	filterd = filterBy(result, func(i *github.Issue) bool {
		return i.CreatedAt.After(dayOfOneYearsAgo)
	})
	showIssue(filterd, "Report less than one year===================================")

	filterd = filterBy(result, func(i *github.Issue) bool {
		return i.CreatedAt.Before(dayOfOneYearsAgo)
	})
	showIssue(filterd, "Report older than one year==================================")
}
