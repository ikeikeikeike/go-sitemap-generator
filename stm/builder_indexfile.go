package stm

func NewBuilderIndexfile(loc *Location) *BuilderIndexfile {
	return &BuilderIndexfile{
		loc:        loc,
	}
}

type BuilderIndexfile struct {
	loc        *Location
}

func (b *BuilderIndexfile) Add(url interface{}) Builder {
	return b
}

func (b *BuilderIndexfile) AddWithErr(url interface{}) (Builder, error) {
	return b, nil
}

func (b *BuilderIndexfile) finalize() { }
func (b *BuilderIndexfile) write() { }
func (b *BuilderIndexfile) run() { }
