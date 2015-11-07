// General sitemap guidelines: https://support.google.com/webmasters/answer/183668
// Number of URLs = 50,000
// File size ( uncompressed ) = 50MB
package stm

const (
	MaxSitemapFiles    = 50000    // max sitemap links per index file
	MaxSitemapLinks    = 50000    // max links per sitemap
	MaxSitemapImages   = 1000     // max images per url
	MaxSitemapNews     = 1000     // max news sitemap per index_file
	MaxSitemapFilesize = 10485760 // bytes
)

const (
	SchemaGeo     = "http://www.google.com/geo/schemas/sitemap/1.0"
	SchemaImage   = "http://www.google.com/schemas/sitemap-image/1.1"
	SchemaMobile  = "http://www.google.com/schemas/sitemap-mobile/1.0"
	SchemaNews    = "http://www.google.com/schemas/sitemap-news/0.9"
	SchemaPagemap = "http://www.google.com/schemas/sitemap-pagemap/1.0"
	SchemaVideo   = "http://www.google.com/schemas/sitemap-video/1.1"
)
