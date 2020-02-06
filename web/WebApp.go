package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type sitemapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

type news struct {
	Titles    []string `xml:"url>news>title"`
	Keywords  []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

type newsMap struct {
	Keyworkd string
	Location string
}

func main() {
	var s sitemapIndex
	var n news
	myNewsMap := make(map[string]newsMap)
	resp, err := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	if err != nil {
		log.Fatal(err)
	}
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	xml.Unmarshal(bytes, &s)

	for _, Location := range s.Locations {
		url := strings.TrimSpace(Location)
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		bytes, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		xml.Unmarshal(bytes, &n)
		for idx := range n.Titles {
			myNewsMap[n.Titles[idx]] = newsMap{n.Keywords[idx], n.Locations[idx]}
		}
	}
	for k, v := range myNewsMap {
		fmt.Println("\n\n\n", k)
		fmt.Println("\n", v.Keyworkd)
		fmt.Println("\n", v.Location)
	}
}
