package stm

import (
	"bytes"
	"time"

	"github.com/beevik/etree"
	"github.com/ikeikeikeike/go-sitemap-generator/stm/utils"
)

func NewSitemapIndexURL(url URL) *sitemapIndexURL {
	return &sitemapIndexURL{data: url}
}

type sitemapIndexURL struct {
	data URL
}

func (su *sitemapIndexURL) XML() []byte {
	doc := etree.NewDocument()
	sitemap := doc.CreateElement("sitemap")

	utils.SetElementValue(sitemap, su.data, "loc")

	if !utils.SetElementValue(sitemap, su.data, "lastmod") {
		lastmod := sitemap.CreateElement("lastmod")
		lastmod.SetText(time.Now().Format(time.RFC3339))
	}

	buf := &bytes.Buffer{}
	// doc.Indent(2)
	doc.WriteTo(buf)

	return buf.Bytes()
}
