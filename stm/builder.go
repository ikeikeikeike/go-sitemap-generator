package stm

type BuilderError interface {
    error
    FullError() bool
    FinalizedError() bool
}

type Builder interface {
	// Content() string
	Add(interface{}) BuilderError
	// AddWithErr(url interface{}) (Builder, error)
	// location() *Location

	Finalize()
	Write()
	run()
}
