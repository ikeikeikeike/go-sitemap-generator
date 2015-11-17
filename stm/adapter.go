package stm

import "regexp"

var GzipPtn = regexp.MustCompile(".gz$")

type Adapter interface {
	Write(loc *Location, data []byte)
}
