package stm

import "log"

// import (
// "sync"
// )

func NewBuilderFile() *BuilderFile {
	return &BuilderFile{
		xmlContent: "",
		write:      make(chan sitemapURL),
		// mu: sync.RWMutex{},
	}
}

type BuilderFile struct {
	xmlContent string // We can use this later
	write      chan sitemapURL
	// mu    sync.RWMutex

	urls []interface{} // XXX: For debug
}

func (b *BuilderFile) Add(url interface{}) Builder {
	smu, err := NewSitemapURL(url)
	if err != nil {
		log.Fatal("Sitemap Key: ", err)
	}
	b.write <- smu; b.urls = append(b.urls, url) // XXX: For debug
	return b
}

func (b *BuilderFile) AddWithErr(url interface{}) (Builder, error) {
	smu, err := NewSitemapURL(url)
	if err != nil {
		log.Fatal("Sitemap Key: ", err)
	}
	b.write <- smu; b.urls = append(b.urls, url) // XXX: For debug
	return b, err
}

func (b *BuilderFile) Content() string {
	return b.xmlContent
}

func (b *BuilderFile) run() {
	for {
		select {
		case sitemapurl := <-b.write:
			b.xmlContent += sitemapurl.Xml() // TODO: Sitemap xml have limit length
			// b.xmlContent += NewSitemapURL(url).Xml() // TODO: Sitemap xml have limit length
		}
	}
}
