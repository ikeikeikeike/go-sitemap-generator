package sitemap

type BuilderFile struct {
	xmlContent string // We can use this later
}

func (b *BuilderFile) Add(url URL) Builder {
	b.xmlContent += NewSitemapURL(url).ToXML()  // TODO: Sitemap xml have limit length
	return b
}
