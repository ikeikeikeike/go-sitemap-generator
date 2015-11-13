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

var (
	IndexXMLHeader = []byte(`<?xml version="1.0" encoding="UTF-8"?>
      <sitemapindex
      xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
      xsi:schemaLocation="http://www.sitemaps.org/schemas/sitemap/0.9
        http://www.sitemaps.org/schemas/sitemap/0.9/siteindex.xsd"
      xmlns="http://www.sitemaps.org/schemas/sitemap/0.9"
    >`)
	IndexXMLFooter = []byte("</sitemapindex>")
)

var (
	XMLHeader = []byte(`<?xml version="1.0" encoding="UTF-8"?>
      <urlset
        xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
        xsi:schemaLocation="http://www.sitemaps.org/schemas/sitemap/0.9
          http://www.sitemaps.org/schemas/sitemap/0.9/sitemap.xsd"
        xmlns="http://www.sitemaps.org/schemas/sitemap/0.9"
        xmlns:image="` + SchemaImage + `"
        xmlns:video="` + SchemaVideo + `"
        xmlns:geo="` + SchemaGeo + `"
        xmlns:news="` + SchemaNews + `"
        xmlns:mobile="` + SchemaMobile + `"
        xmlns:pagemap="` + SchemaPagemap + `"
        xmlns:xhtml="http://www.w3.org/1999/xhtml"
    >`)
	XMLFooter = []byte("</urlset>")
)
