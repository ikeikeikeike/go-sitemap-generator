package sitemap

import "time"

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

type url struct {
    ServerName string `xml:"serverName"`
    ServerIP   string `xml:"serverIP"`
}

func NewSitemapURL(url URL) sitemapURL {
	smu := sitemapURL{url: url}
	return smu
}

type sitemapURL struct {
	url URL
}

func (smu sitemapURL) ToXML() string {
	xml := url{}
	smu.url
	return ""
}
