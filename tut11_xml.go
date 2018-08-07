package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
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

func tut11() {
	var s SitemapIndex
	var n News
	var newsMap sync.Map

	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &s)

	fmt.Printf("Gonna make %v requests", len(s.Locations))
	time1 := time.Now().UnixNano() / 1000000
	for _, location := range s.Locations {
		wg.Add(1)
		go requestAndMap(location, &n, newsMap)
		/* fmt.Printf("\n %v, Titles: %s", index+1, n.Titles) */
	}
	wg.Wait()
	time2 := time.Now().UnixNano() / 1000000
	fmt.Printf("\n%v\n", time2-time1)
	/* for index, value := range newsMap {
		fmt.Println("\n\n\n", index)
		fmt.Printf("\n%s", value.Keyword)
		fmt.Printf("\n%s", value.Location)
	} */
}

func requestAndMap(location string, n *News, newsMap sync.Map) {
	resp, _ := http.Get(location)
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &n)
	for index := range n.Titles {
		/* fmt.Println("Hello", index) */
		newsMap.Store(n.Titles[index], NewsMap{n.Keywords[index], n.Locations[index]})
	}
	wg.Done()
}

/* package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

func tut11() {
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	var s SitemapIndex
	xml.Unmarshal(bytes, &s)
	fmt.Printf("Here %s some %s", "are", "variables")
	for index, location := range s.Locations {
		fmt.Printf("\n %v, %s", index+1, location)
	}
} */
