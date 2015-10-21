package stm

import (
	"fmt"
	"log"
)

// import (
// "sync"
// )

func NewBuilderFile(loc *Location) *BuilderFile {
	return &BuilderFile{
		xmlContent: "",
		write:      make(chan sitemapURL),
		loc:        loc,
		// mu: sync.RWMutex{},
	}
}

type BuilderFile struct {
	xmlContent string // We can use this later
	write      chan sitemapURL
	loc        *Location

	urls []interface{} // XXX: For debug
}

func (b *BuilderFile) Add(url interface{}) Builder {
	smu, err := NewSitemapURL(url)
	if err != nil {
		panic(fmt.Sprintf("[F] Sitemap: %s", err))
	}
	b.xmlContent += smu.Xml() // TODO: Sitemap xml have limit length
	// b.write <- smu; b.urls = append(b.urls, url) // XXX: For debug
	return b
}

func (b *BuilderFile) AddWithErr(url interface{}) (Builder, error) {
	smu, err := NewSitemapURL(url)
	if err != nil {
		log.Println("[E] Sitemap: ", err)
	}
	b.xmlContent += smu.Xml() // TODO: Sitemap xml have limit length
	// b.write <- smu; b.urls = append(b.urls, url) // XXX: For debug
	return b, err
}

func (b *BuilderFile) Content() string {
	return b.xmlContent
}

// func (b *BuilderFile) location() *Location {
	// return b.loc
// }

func (b *BuilderFile) run() {
	for {
		select {
		case smu := <-b.write:
			b.xmlContent += smu.Xml() // TODO: Sitemap xml have limit length
		}
	}
}
