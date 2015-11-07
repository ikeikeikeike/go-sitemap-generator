package stm

import "github.com/k0kubun/pp"

func NewBuilderIndexfile(loc *Location) *BuilderIndexfile {
	return &BuilderIndexfile{
		loc: loc,
	}
}

type BuilderIndexfile struct {
	loc      *Location
	linkcnt  int
	totalcnt int
}

func (b *BuilderIndexfile) Add(link interface{}) BuilderError {
	bldr := link.(*BuilderFile)

	b.totalcnt += bldr.linkcnt

	if !bldr.isFinalized() {
		bldr.Finalize()
	}

	// TODO: first sitemap
	// if b.linkcnt == 0 { }

	bldr.Write()
	return nil
}

// func (b *BuilderIndexfile) AddWithErr(url interface{}) (Builder, error) {
// return b, nil
// }

func (b *BuilderIndexfile) Finalize() {}
func (b *BuilderIndexfile) Write() {
	pp.Println("write indexfile")
}

func (b *BuilderIndexfile) run() {}
