package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"./fileio"
	"./myhttp"
)

const api = "http://www.omdbapi.com/"

// Rating :Ratingフィールド
type Rating struct {
	Source string
	Value  string
}

// Info :映画データ
type Info struct {
	Title      string
	Year       string
	Rated      string
	Released   string
	Runtime    string
	Genre      string
	Director   string
	Writer     string
	Actor      string
	Plot       string
	Language   string
	Country    string
	Awards     string
	Poster     string
	Ratings    []Rating
	Metascore  string
	imdbRating string
	imdbVotes  string
	imdbID     string
	Type       string
	DVD        string
	BoxOffice  string
	Production string
	Website    string
	Response   string
}

func main() {
	if len(os.Args) < 2 {
		return
	}

	// JSONデータの取得
	query := url.QueryEscape(strings.Join(os.Args[1:], " "))
	var (
		data []byte
		err  error
	)
	address := api + "?t=" + query
	if data, err = myhttp.Get(address); err != nil {
		log.Fatal(err)
	}

	// JSONのデコード
	var info Info
	if err = json.NewDecoder(bytes.NewBuffer(data)).Decode(&info); err != nil {
		log.Fatal(err)
	}

	// 該当する映画は存在しない
	if info.Response == "False" {
		log.Fatal(fmt.Errorf("%s:not found", query))
	}

	// posterの取得
	if data, err = myhttp.Get(info.Poster); err != nil {
		log.Fatal(err)
	}

	//posterの書き込み
	if err = fileio.WriteFile(info.Title+".jpg", data); err != nil {
		log.Fatal(err)
	}
}
