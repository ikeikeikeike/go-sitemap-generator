package stm

import (
	"fmt"
	"log"
)

func NewNamer(base string) *Namer {
	namer := &Namer{
		base: base,
		opts: options{
			zero:      0,
			extension: ".xml.gz",
			start:     1,
		},
	}
	namer.Reset()
	return namer
}

type options struct {
	zero      int
	extension string
	start     int
}

type Namer struct {
	base  string
	count int
	opts  options
}

func (n *Namer) String() string {
	ext := n.opts.extension
	return fmt.Sprintf("%s%d%s", n.base, n.count, ext)
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
		log.Fatal("Already at the start of the series")
	}
	if n.count <= n.opts.start {
		n.count = n.opts.zero
	} else {
		n.count -= 1
	}
	return n
}
