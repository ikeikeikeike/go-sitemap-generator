package stm

type BuilderError interface {
    error
    FullError() bool
}

type Builder interface {
	Content() []byte
	Add(interface{}) BuilderError
	Write()
	run()
}

type URL map[string]interface{}
