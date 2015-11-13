package stm

import "bytes"

func NewBuilderIndexfile(loc *Location) *BuilderIndexfile {
	return &BuilderIndexfile{
		loc: loc,
	}
}

type BuilderIndexfile struct {
	loc      *Location
	content  []byte
	linkcnt  int
	totalcnt int
}

func (b *BuilderIndexfile) Add(link interface{}) BuilderError {
	bldr := link.(*BuilderFile)
	bldr.Write()

	smu := NewSitemapIndexURL(URL{"loc": bldr.loc.URL()})
	b.content = append(b.content, smu.XML()...)

	b.totalcnt += bldr.linkcnt
	b.linkcnt += 1
	return nil
}

func (b *BuilderIndexfile) Content() []byte {
	return b.content
}

func (b *BuilderIndexfile) Write() {
	c := bytes.Join(bytes.Fields(IndexXMLHeader), []byte(" "))
	c = append(append(c, b.Content()...), IndexXMLFooter...)

	b.loc.Write(c, b.linkcnt)
}

func (b *BuilderIndexfile) run() {}
