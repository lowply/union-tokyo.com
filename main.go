package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"net/http"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

type Article struct {
	Category    string
	Title       string
	Description string
	Date        string
	Content     string
}

type Artist struct {
	Article
	Name       string
	Alias      string
	Atype      string
	Org        string
	Suborg     string
	Profile_ja string
	Profile_en string
	Links      []Link
}

type Link struct {
	Name string
	Url  string
}

type Interview struct {
	Artist
	QAs []QA
}

type QA struct {
	Question string
	Answer   string
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func main() {
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		Layout: "layout",
	}))
	martini.Env = martini.Prod
	m.Get("/", top)
	m.Get("/about", about)
	m.Get("/column", column)
	m.Get("/column/interview", column_interview)
	m.Get("/column/interview/:name", column_interview_artist)
	m.Get("/artist", artist_all)
	m.Get("/artist/:name", artist)
	m.Get("/time", time)
	m.Get("/report", report)
	http.ListenAndServe("127.0.0.1:3000", m)
	m.Run()
}

func top(r render.Render) {
	article := Article{
		Category:    "top",
		Title:       "\"union\" OF TOKYO DRUM&BASS, AND CHARITY PARTY",
		Description: "史上初！東京ドラムンベース パーティーが集結するスペシャル パーティーが、チャリティー パーティーで開催！",
		Date:        "22 JUL 2014 TUE",
	}
	r.HTML(200, "top", article)
}

func about(r render.Render) {
	article := Article{
		Category:    "about",
		Title:       "東京ドラムンベース  シーンが、遂にひとつとなる union(集結) !!",
		Description: "東京ドラムンベース パーティーが集結する union とは!?",
		Date:        "22 JUL 2014 TUE",
	}
	r.HTML(200, "about", article)
}

func column(r render.Render) {
	article := Article{
		Category:    "column",
		Title:       "union COLUMN",
		Description: "union COLUMN",
	}
	r.HTML(200, "column", article)
}

func artist_all(r render.Render) {
	article := Article{
		"artist",
		"union LINEUP",
		"LINEUP - union 2014.08.01 Friday @ Glad + LOUNGE NEO",
		"",
		"",
	}
	r.HTML(200, "artist_all", article)
}

func column_interview(r render.Render) {
	article := Article{
		Category:    "column",
		Title:       "union ARTIST INTERVIEW",
		Description: "union OF TOKYO DRUM&BASS 出演者へ、それぞれのドラムンベースについてのインタビュー！",
	}
	r.HTML(200, "column_interview", article)
}

func time(r render.Render) {
	article := Article{
		Category:    "time",
		Title:       "TIME TABLE",
		Description: "TIME TABLE - \"union\" 2014.08.01 Friday @ Glad + LOUNGE NEO",
	}
	r.HTML(200, "time", article)
}

func report(r render.Render) {
	article := Article{
		Category:    "report",
		Title:       "8月1日 (金) 、\"union\" of Tokyo D&B の寄付に関するご報告。",
		Description: "REPORT - \"union\" OF TOKYO DRUM&BASS",
	}
	r.HTML(200, "report", article)
}

func column_interview_artist(params martini.Params, r render.Render) {
	name := params["name"]
	file, err := ioutil.ReadFile("data/artists.json")
	if err != nil {
		log.Fatal(err)
	}
	var artists map[string]Artist
	err = json.Unmarshal(file, &artists)
	if err != nil {
		log.Fatal(err)
	}

	file, err = ioutil.ReadFile("data/interviews.json")
	if err != nil {
		log.Fatal(err)
	}
	var interviews map[string]*Interview
	err = json.Unmarshal(file, &interviews)
	if err != nil {
		log.Fatal(err)
	}
	t := "INTERVIEW - " + artists[name].Name + " (" + artists[name].Org + ")"
	article := Article{
		Category:    "column",
		Title:       t,
		Description: t,
		Date:        "",
	}

	artist := artists[name]
	interviews[name].Artist = artist
	interviews[name].Artist.Article = article

	// check if param "name" exists
	if _, ok := interviews[name]; ok {
		r.HTML(200, "interview", interviews[name])
	}
}

func artist(params martini.Params, r render.Render) {
	name := params["name"]
	file, err := ioutil.ReadFile("data/artists.json")
	if err != nil {
		log.Fatal(err)
	}
	var artists map[string]*Artist
	err = json.Unmarshal(file, &artists)
	if err != nil {
		log.Fatal(err)
	}
	t := artists[name].Atype + " - " + artists[name].Name + " (" + artists[name].Org + ")"
	article := Article{
		Category:    "artist",
		Title:       t,
		Description: t,
		Date:        "",
	}
	artists[name].Article = article
	// check if param "name" exists
	if _, ok := artists[name]; ok {
		r.HTML(200, "artist", artists[name])
	}
}
