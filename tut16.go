package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

var wg sync.WaitGroup

type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Titles    []string `xml:"url>news>title"`
	Keywords  []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

type NewsMap struct {
	Keyword  string
	Location string
}

type NewsAggPage struct {
	Title string
	News  map[string]NewsMap
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1> Whoa this is neat!</h1>")
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	var s SitemapIndex
	/* var newsMap sync.Map */
	newsMap := make(map[string]NewsMap)

	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &s)
	resp.Body.Close()
	queue := make(chan News, 22)

	fmt.Printf("Gonna make %v requests", len(s.Locations))
	time1 := time.Now().UnixNano() / 1000000
	for _, location := range s.Locations {
		wg.Add(1)
		go newsRoutine(queue, location)
		/* fmt.Printf("\n %v, Titles: %s", index+1, n.Titles) */
	}
	wg.Wait()
	close(queue)
	for elem := range queue {
		for index := range elem.Keywords {
			newsMap[elem.Titles[index]] = NewsMap{elem.Keywords[index], elem.Locations[index]}
		}
	}
	time2 := time.Now().UnixNano() / 1000000
	fmt.Printf("\n%v\n", time2-time1)
	/* fmt.Println(newsMap) */
	p := NewsAggPage{Title: "Amazing news agg", News: newsMap}
	t, _ := template.ParseFiles("basicTemplating.html")
	t.Execute(w, p)
	/* fmt.Fprintf(w, "im done") */
}

func newsRoutine(c chan News, location string) {
	defer cleanUp()
	var n News
	resp, err := http.Get(location)
	if err != nil {
		panic("Omg something went wrong!")
	}
	bytes, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		panic("Omg something went wrong2!")
	}
	xml.Unmarshal(bytes, &n)
	resp.Body.Close()
	c <- n
}

func tut16() {
	r := mux.NewRouter()

	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/agg/", newsAggHandler)
	log.Fatal(http.ListenAndServe(":8081", r))
}
