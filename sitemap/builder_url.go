package sitemap

import (
	"bytes"
	"fmt"
	"time"

	"github.com/beevik/etree"
	"github.com/kr/pretty"
)

type URL struct {
	Priority   float32
	Changefreq string
	Lastmod    time.Time
	Expires    time.Time
	Host       string
	Loc        string
	Images     string
	Geo        string
	Mobile     bool
	Alternates string
	Pagemap    string
}

func NewSitemapURL(url URL) sitemapURL {
	su := sitemapURL{url: url}
	return su
}

type sitemapURL struct {
	url URL
}

func (su sitemapURL) ToXML() string {
	doc := etree.NewDocument()
	url := doc.CreateElement("url")

	if su.url.Priority > 0 {
		priority := url.CreateElement("priority")
		priority.SetText(fmt.Sprint("%f", su.url.Priority))
	}

	if su.url.Changefreq != "" {
		changefreq := url.CreateElement("changefreq")
		changefreq.SetText(su.url.Changefreq)
	}

	if su.url.Mobile {
		_ = url.CreateElement("mobile:mobile")
	}

	buf := &bytes.Buffer{}
	doc.Indent(2)
	doc.WriteTo(buf)

	st := buf.String()
	pretty.Println(st)
	println("")

	return st
}
