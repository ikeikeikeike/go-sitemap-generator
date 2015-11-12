package stm

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

	smu := NewSitemapIndexURL(URL{"loc": bldr.loc.Filename()})
	b.content = append(b.content, smu.Xml()...)

	b.totalcnt += bldr.linkcnt
	b.linkcnt += 1
	return nil
}

func (b *BuilderIndexfile) Content() []byte {
	return b.content
}

func (b *BuilderIndexfile) Write() {
	// TODO: header and footer
	// TODO: Change loc.Filename, loc.Path
	b.loc.Write(b.Content(), b.linkcnt) // @location.write(@xml_wrapper_start + @xml_content + @xml_wrapper_end, link_count)
}

func (b *BuilderIndexfile) run() {}
