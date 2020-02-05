package main

/*
<sitemapindex xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
<sitemap>
<loc>
https://www.washingtonpost.com/news-sitemaps/politics.xml
</loc>
</sitemap>
<sitemap>
<loc>
https://www.washingtonpost.com/news-sitemaps/opinions.xml
</loc>
</sitemap>
<sitemap>
<loc>
https://www.washingtonpost.com/news-sitemaps/local.xml
</loc>
</sitemap>
<sitemap>
<loc>
https://www.washingtonpost.com/news-sitemaps/sports.xml
</loc>
</sitemap>
<sitemap>
<loc>
https://www.washingtonpost.com/news-sitemaps/national.xml
</loc>
</sitemap>
<sitemap>
<loc>
https://www.washingtonpost.com/news-sitemaps/world.xml
</loc>
</sitemap>
<sitemap>
<loc>
https://www.washingtonpost.com/news-sitemaps/business.xml
</loc>
</sitemap>
<sitemap>
<loc>
https://www.washingtonpost.com/news-sitemaps/technology.xml
</loc>
</sitemap>
<sitemap>
<loc>
https://www.washingtonpost.com/news-sitemaps/lifestyle.xml
</loc>
</sitemap>
<sitemap>
<loc>
https://www.washingtonpost.com/news-sitemaps/entertainment.xml
</loc>
</sitemap>
<sitemap>
<loc>
https://www.washingtonpost.com/news-sitemaps/goingoutguide.xml
</loc>
</sitemap>
</sitemapindex>
*/
import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type SitemapIndex struct {
	Locations []Location `xml:"sitemap"`
}

type Location struct {
	Loc string `xml:"loc"`
}

func (l Location) String() string {
	//return fmt.Sprintf(l.Loc)
	return l.Loc
}

func main() {
	resp, err := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	if err != nil {
		log.Fatal(err)
	}
	bytes, _ := ioutil.ReadAll(resp.Body)

	resp.Body.Close()
	var s SitemapIndex
	xml.Unmarshal(bytes, &s)
	fmt.Println(s.Locations)
}
