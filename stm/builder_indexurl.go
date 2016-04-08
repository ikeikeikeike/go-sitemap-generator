package stm

import (
	"time"

	"github.com/beevik/etree"
)

// NewSitemapIndexURL and NewSitemapURL are almost the same behavior.
func NewSitemapIndexURL(url URL) SitemapURL {
	return &sitemapIndexURL{data: url}
}

// sitemapIndexURL and sitemapURL are almost the same behavior.
type sitemapIndexURL struct {
	data URL
}

// XML and sitemapIndexURL.XML are almost the same behavior.
func (su *sitemapIndexURL) XML() []byte {
	doc := etree.NewDocument()
	sitemap := doc.CreateElement("sitemap")

	SetBuilderElementValue(sitemap, su.data, "loc")

	if !SetBuilderElementValue(sitemap, su.data, "lastmod") {
		lastmod := sitemap.CreateElement("lastmod")
		lastmod.SetText(time.Now().Format(time.RFC3339))
	}

	buf := poolBuffer.Get()
	// doc.Indent(2)
	doc.WriteTo(buf)

	bytes := buf.Bytes()
	poolBuffer.Put(buf)

	return bytes
}
