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

type BuilderURL struct {
	// TODO: Its change to struct coz sitemap xml have limit length
	//		 and that append is slowly runnning.
	urls []URL
}

func (b *BuilderURL) Add(url URL) Builder {
	b.urls = append(b.urls, url)
	return b
}
