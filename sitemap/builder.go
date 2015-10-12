package sitemap

type Builder interface {
	Add(URL) Builder
}
