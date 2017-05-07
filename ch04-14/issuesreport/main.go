package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"./format"
	"./github"
)

var logger *log.Logger

func init() {
	logger = log.New(os.Stderr, "issuesreport: ", log.LstdFlags)
}

func main() {
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/search", handleSearch)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	// 入力フォームなどは表示せず,直接アドレスバーへ入力.
	body := "<h2>example request:<br/>localhost:8000/search?q=keyword1+keyword2+...</h2>"
	w.Write([]byte(body))
}
func handleSearch(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm() // Getの前にParseFormしておく必要あり
	v := r.Form.Get("q")
	// Request.From.Get は空白でつながった単一の文字列を返すので
	// 本来ならurl decode処理が必要と思う.
	// 今回は受付時のencodeもやってないのでそのまま.
	keywords := strings.Split(v, " ")
	result, err := github.SearchIssues(keywords)
	if err != nil {
		logger.Print(err)
		w.WriteHeader(400) // Bad Request
		body := "<h1>Error! Bad Request</h1>"
		w.Write([]byte(body))
	} else {
		w.Header().Set("Content-Type", "text/html")
		format.OutputHTML(w, result)
	}
}
