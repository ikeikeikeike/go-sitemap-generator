package stm

import "log"

type builderFileError struct {
	error
	full      bool
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
	content      []byte
	build        chan sitemapURL
	loc          *Location
	linkcnt      int
	newscnt      int

	urls []interface{} // XXX: For debug
}

func (b *BuilderFile) Add(url interface{}) BuilderError {
	smu, err := NewSitemapURL(url)
	if err != nil {
		log.Fatalln("[F] Sitemap: %s", err)
	}

	bytes := smu.Xml()

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
	b.content = make([]byte, MaxSitemapLinks, MaxSitemapFilesize)
}

func (b *BuilderFile) Content() []byte {
	return b.content
}

func (b *BuilderFile) Write() {
	b.loc.ReserveName()

	// TODO: header and footer
	b.loc.Write(b.Content(), b.linkcnt) // @location.write(@xml_wrapper_start + @xml_content + @xml_wrapper_end, link_count)

	b.clear() // @xml_content = @xml_wrapper_start = @xml_wrapper_end = ''
}

func (b *BuilderFile) run() {
	for {
		select {
		case smu := <-b.build:
			b.content = append(b.content, smu.Xml()...) // TODO: Sitemap xml have limit length
		}
	}
}
