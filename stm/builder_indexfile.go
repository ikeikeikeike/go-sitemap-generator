package stm

func NewBuilderIndexfile(loc *Location) *BuilderIndexfile {
	return &BuilderIndexfile{
		loc: loc,
	}
}

type BuilderIndexfile struct {
	loc           *Location
	linkCount     int
	bldrLinkCount int
}

func (b *BuilderIndexfile) Add(link interface{}) Builder {
	bldr := link.(Builder)
	bldr.write()
	return bldr
}

func (b *BuilderIndexfile) AddWithErr(url interface{}) (Builder, error) {
	return b, nil
}

func (b *BuilderIndexfile) isFull() bool {
	return true
}

func (b *BuilderIndexfile) isFinalized() bool {
	return true
}

func (b *BuilderIndexfile) finalize() {}
func (b *BuilderIndexfile) write()    {}
func (b *BuilderIndexfile) run()      {}
