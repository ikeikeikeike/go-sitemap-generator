package stm

type Builder interface {
	Content() string
	Add(interface{}) Builder
	run()
}
