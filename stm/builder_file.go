package stm

import (
	"fmt"
	"log"
)

func NewBuilderFile(loc *Location) *BuilderFile {
	return &BuilderFile{
		xmlContent: "",
		build:      make(chan sitemapURL),
		loc:        loc,
	}
}

type BuilderFile struct {
	xmlContent string // We can use this later
	build      chan sitemapURL
	loc        *Location

	urls []interface{} // XXX: For debug
}

func (b *BuilderFile) Add(url interface{}) Builder {
	smu, err := NewSitemapURL(url)
	if err != nil {
		panic(fmt.Sprintf("[F] Sitemap: %s", err))
	}
	b.xmlContent += smu.Xml() // TODO: Sitemap xml have limit length
	// b.build <- smu; b.urls = append(b.urls, url) // XXX: For debug
	return b
}

func (b *BuilderFile) AddWithErr(url interface{}) (Builder, error) {
	smu, err := NewSitemapURL(url)
	if err != nil {
		log.Println("[E] Sitemap: ", err)
	}
	b.xmlContent += smu.Xml() // TODO: Sitemap xml have limit length
	// b.build <- smu; b.urls = append(b.urls, url) // XXX: For debug
	return b, err
}

func (b *BuilderFile) Content() string {
	return b.xmlContent
}

// func (b *BuilderFile) location() *Location {
// return b.loc
// }

func (b *BuilderFile) finalize() {}
func (b *BuilderFile) write()    {
	
        // raise SitemapGenerator::SitemapError.new("Sitemap already written!") if written?
        // finalize! unless finalized?
        // reserve_name
        // @location.write(@xml_wrapper_start + @xml_content + @xml_wrapper_end, link_count)
        // @xml_content = @xml_wrapper_start = @xml_wrapper_end = ''
        // @written = true
}

func (b *BuilderFile) run() {
	for {
		select {
		case smu := <-b.build:
			b.xmlContent += smu.Xml() // TODO: Sitemap xml have limit length
		}
	}
}
