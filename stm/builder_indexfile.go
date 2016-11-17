package stm

import "bytes"

// NewBuilderIndexfile returns the created the BuilderIndexfile's pointer
func NewBuilderIndexfile(loc *Location) *BuilderIndexfile {
	return &BuilderIndexfile{loc: loc}
}

// BuilderIndexfile provides implementation for the Builder interface.
type BuilderIndexfile struct {
	loc      *Location
	content  []byte
	linkcnt  int
	totalcnt int
}

// Add method joins old bytes with creates bytes by it calls from Sitemap.Finalize method.
func (b *BuilderIndexfile) Add(link interface{}) BuilderError {
	bldr := link.(*BuilderFile)
	bldr.Write()

	smu := NewSitemapIndexURL(URL{"loc": bldr.loc.URL()})
	b.content = append(b.content, smu.XML()...)

	b.totalcnt += bldr.linkcnt
	b.linkcnt++
	return nil
}

func (b *BuilderIndexfile) AddLocation(loc *Location) BuilderError {
	smu := NewSitemapIndexURL(URL{"loc": loc.URL()})
	b.content = append(b.content, smu.XML()...)

	b.totalcnt ++
	b.linkcnt++
	return nil
}

// Content and BuilderFile.Content are almost the same behavior.
func (b *BuilderIndexfile) Content() []byte {
	return b.content
}

// XMLContent and BuilderFile.XMLContent share almost the same behavior.
func (b *BuilderIndexfile) XMLContent() []byte {
	c := bytes.Join(bytes.Fields(IndexXMLHeader), []byte(" "))
	c = append(append(c, b.Content()...), IndexXMLFooter...)

	return c
}

// clear will initialize xml content.
func (b *BuilderIndexfile) clear() {
	b.content = make([]byte, 0, MaxSitemapFilesize)
}

// Write and Builderfile.Write are almost the same behavior.
func (b *BuilderIndexfile) Write() {
	b.loc.ReserveName()
	c := b.XMLContent()

	b.loc.Write(c, b.linkcnt)
	b.clear()
}
