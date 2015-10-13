package sitemap

import (
	"sync"
)

func NewBuilderFile() *BuilderFile {
	return &BuilderFile{
		xmlContent: "",
		write: make(chan sitemapURL),
		mu: sync.RWMutex{},
	}
}

type BuilderFile struct {
	xmlContent string // We can use this later

	write chan sitemapURL
	mu    sync.RWMutex

	urls []URL // For debug
}

func (b *BuilderFile) Add(url URL) Builder {
	// b.urls = append(b.urls, url)               // For debug

	sitemapurl := NewSitemapURL(url)
	b.write <- sitemapurl
	return b
}

func (b *BuilderFile) run() {
	for {
		select {
		case sitemapurl := <-b.write:
			b.xmlContent += sitemapurl.ToXML() // TODO: Sitemap xml have limit length

			// cmd.result <- ldb.execGet(cmd)
			// case <-updateTick.C:
			// ldb.mu.RLock()
			// if !ldb.downloading {
			// go ldb.download(ctx, true)
			// }
			// ldb.mu.RUnlock()
		}
	}
}
