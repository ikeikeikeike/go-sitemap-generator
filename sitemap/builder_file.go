package sitemap

type BuilderFile struct {
	xmlContent string // We can use this later
	urls       []URL  // For debug
}

func (b *BuilderFile) Add(url URL) Builder {
	b.urls = append(b.urls, url)               // For debug
	b.xmlContent += NewSitemapURL(url).ToXML() // TODO: Sitemap xml have limit length
	return b
}
