package stm

type Adapter interface {
	Write(loc *Location, data []byte)
}
