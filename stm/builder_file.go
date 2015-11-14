package stm

import (
	"bytes"
	"log"
)

type builderFileError struct {
	error
	full bool
}

func (e *builderFileError) FullError() bool {
	return e.full
}

func NewBuilderFile(loc *Location) *BuilderFile {
	b := &BuilderFile{
		build: make(chan sitemapURL),
		loc:   loc,
	}
	b.clear()
	return b
}

type BuilderFile struct {
	content []byte
	build   chan sitemapURL
	loc     *Location
	linkcnt int
	newscnt int

	urls []interface{} // XXX: For debug
}

func (b *BuilderFile) Add(url interface{}) BuilderError {
	smu, err := NewSitemapURL(url)
	if err != nil {
		log.Fatalf("[F] Sitemap: %s", err)
	}

	bytes := smu.XML()

	if !b.isFileCanFit(bytes) {
		return &builderFileError{error: err, full: true}
	}

	b.content = append(b.content, bytes...) // TODO: Sitemap xml have limit length
	b.linkcnt += 1
	// b.build <- smu; b.urls = append(b.urls, url) // XXX: For debug
	return nil
}

func (b *BuilderFile) isFileCanFit(bytes []byte) bool {
	r := len(append(b.content, bytes...)) < MaxSitemapFilesize
	r = r && b.linkcnt < MaxSitemapLinks
	return r && b.newscnt < MaxSitemapNews
}

func (b *BuilderFile) clear() {
	// b.content = make([]byte, MaxSitemapLinks, MaxSitemapFilesize)
	b.content = make([]byte, 0, MaxSitemapFilesize)
}

func (b *BuilderFile) Content() []byte {
	return b.content
}

func (b *BuilderFile) Write() {
	b.loc.ReserveName()

	c := bytes.Join(bytes.Fields(XMLHeader), []byte(" "))
	c = append(append(c, b.Content()...), XMLFooter...)

	b.loc.Write(c, b.linkcnt)
	b.clear() // @xml_content = @xml_wrapper_start = @xml_wrapper_end = ''
}

func (b *BuilderFile) run() {
	for {
		select {
		case smu := <-b.build:
			b.content = append(b.content, smu.XML()...) // TODO: Sitemap xml have limit length
		}
	}
}
