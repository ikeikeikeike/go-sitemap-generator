package stm

import (
	"fmt"
	"log"
)

func NewNamer(opts *NOpts) *Namer {
	if opts.extension == "" {
		opts.extension = ".xml.gz"
	}

	namer := &Namer{opts: opts}
	namer.Reset()
	return namer
}

type NOpts struct {
	base      string
	zero      int
	extension string
	start     int
}

type Namer struct {
	count int
	opts  *NOpts
}

func (n *Namer) String() string {
	ext := n.opts.extension
	if n.count == 0 {
		return fmt.Sprintf("%s%s", n.opts.base, ext)
	}
	return fmt.Sprintf("%s%d%s", n.opts.base, n.count, ext)
}

func (n *Namer) Reset() {
	n.count = n.opts.zero
}

func (n *Namer) IsStart() bool {
	return n.count == n.opts.zero
}

func (n *Namer) Next() *Namer {
	if n.IsStart() {
		n.count = n.opts.start
	} else {
		n.count += 1
	}
	return n
}

func (n *Namer) Previous() *Namer {
	if n.IsStart() {
		log.Fatal("[F] Already at the start of the series")
	}
	if n.count <= n.opts.start {
		n.count = n.opts.zero
	} else {
		n.count -= 1
	}
	return n
}
