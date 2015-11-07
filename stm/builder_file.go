package stm

import "log"

type builderFileError struct {
	error
	full      bool
	finalized bool
}

func (e *builderFileError) FullError() bool {
	return e.full
}

func (e *builderFileError) FinalizedError() bool {
	return e.finalized
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
	frozen       bool
	linkcnt      int
	newscnt      int
	written      bool
	reservedName string

	urls []interface{} // XXX: For debug
}

func (b *BuilderFile) Add(url interface{}) BuilderError {
	smu, err := NewSitemapURL(url)
	if err != nil {
		log.Fatalln("[F] Sitemap: %s", err)
	}

	bytes := smu.Xml()

	if b.isFinalized() {
		return &builderFileError{error: err, finalized: true}
	} else if !b.isFileCanFit(bytes) {
		return &builderFileError{error: err, full: true}
	}

	// TODO: News sitemap xml
	// if smu.isNews() {
	// b.newscnt += 1
	// }

	b.content = append(b.content, bytes...) // TODO: Sitemap xml have limit length
	b.linkcnt += 1
	// b.build <- smu; b.urls = append(b.urls, url) // XXX: For debug
	return nil
}

// func (b *BuilderFile) AddWithErr(url interface{}) (Builder, error) {
// smu, err := NewSitemapURL(url)
// if err != nil {
// log.Println("[E] Sitemap: ", err)
// }
// b.content += smu.Xml() // TODO: Sitemap xml have limit length
// // b.build <- smu; b.urls = append(b.urls, url) // XXX: For debug
// return b, nil
// }

func (b *BuilderFile) Content() []byte {
	return b.content
}

func (b *BuilderFile) Finalize() {
	b.frozen = true
}

func (b *BuilderFile) isFinalized() bool {
	return b.frozen
}

func (b *BuilderFile) isWritten() bool {
	return b.written
}

func (b *BuilderFile) isFileCanFit(bytes []byte) bool {
	r := len(append(b.content, bytes...)) < MaxSitemapFilesize
	r = r && b.linkcnt < MaxSitemapLinks
	return r && b.newscnt < MaxSitemapNews
}

// func (b *BuilderFile) location() *Location {
// return b.loc
// }

func (b *BuilderFile) setReverseName() {
	if b.reservedName == "" {
		b.reservedName = b.loc.ReserveName()
	}
}

func (b *BuilderFile) clear() {
	b.content = make([]byte, MaxSitemapLinks, MaxSitemapFilesize)
}

func (b *BuilderFile) Write() {
	if b.isWritten() {
		log.Fatalln("[F] Sitemap already written!")
	}

	if !b.isFinalized() {
		b.Finalize()
	}

	b.setReverseName()

	// TODO: header and footer
	b.loc.Write(b.content, b.linkcnt) // @location.write(@xml_wrapper_start + @xml_content + @xml_wrapper_end, link_count)

	b.clear() // @xml_content = @xml_wrapper_start = @xml_wrapper_end = ''
	b.written = true
}

func (b *BuilderFile) run() {
	for {
		select {
		case smu := <-b.build:
			b.content = append(b.content, smu.Xml()...) // TODO: Sitemap xml have limit length
		}
	}
}
