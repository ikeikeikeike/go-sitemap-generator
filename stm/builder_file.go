package stm

// import (
	// "sync"
// )

func NewBuilderFile() *BuilderFile {
	return &BuilderFile{
		xmlContent: "",
		write: make(chan sitemapURL),
		// mu: sync.RWMutex{},
	}
}

type BuilderFile struct {
	xmlContent string // We can use this later
	write chan sitemapURL
	// mu    sync.RWMutex

	urls []URL // For debug
}

func (b *BuilderFile) Add(url interface{}) Builder {
	// b.xmlContent += NewSitemapURL(url).Xml() // TODO: Sitemap xml have limit length
	b.write <- NewSitemapURL(url)
	return b
}

func (b *BuilderFile) Content() string {
	return b.xmlContent
}

func (b *BuilderFile) run() {
	for {
		select {
		case sitemapurl := <-b.write:
			b.xmlContent += sitemapurl.Xml() // TODO: Sitemap xml have limit length
		}
	}
}
