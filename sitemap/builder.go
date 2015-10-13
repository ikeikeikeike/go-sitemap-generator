package sitemap

type Builder interface {
	Content() string
	Add(URL) Builder
	run()
}
