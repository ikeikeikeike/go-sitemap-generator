package stm

import (
	"bytes"
	"fmt"
	"time"

	"github.com/beevik/etree"
)

func NewSitemapIndexURL(url interface{}) *sitemapIndexURL {
	return &sitemapIndexURL{data: url.(URL)}
}

type sitemapIndexURL struct {
	data URL
}

func (su *sitemapIndexURL) Xml() []byte {
	doc := etree.NewDocument()
	sitemap := doc.CreateElement("sitemap")

	if v, ok := su.data["loc"]; ok {
		loc := sitemap.CreateElement("loc")
		loc.SetText(fmt.Sprint(v))
	}

	lastmod := sitemap.CreateElement("lastmod")
	lastmod.SetText(time.Now().Format(time.RFC3339))

	buf := &bytes.Buffer{}
	// doc.Indent(2)
	doc.WriteTo(buf)

	return buf.Bytes()
}
