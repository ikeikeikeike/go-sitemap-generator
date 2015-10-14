package stm

import (
	"bytes"
	"fmt"
	"time"

	"github.com/beevik/etree"
)

// http://www.sitemaps.org/protocol.html
// https://support.google.com/webmasters/answer/178636
type URLModel struct {
	Priority   float32                `valid:"float,length(0.0|1.0)"`
	Changefreq string                 `valid:"alpha(always|hourly|daily|weekly|monthly|yearly|never)"`
	Lastmod    time.Time              `valid:"-"`
	Expires    time.Time              `valid:"-"`
	Host       string                 `valid:"ipv4"`
	Loc        string                 `valid:"url"`
	Images     string                 `valid:"url"`
	Geo        string                 `valid:"latitude,"longitude`
	Mobile     bool                   `valid:"-"`
	Alternates map[string]interface{} `valid:"-"`
	Pagemap    map[string]interface{} `valid:"-"`
}

type URL map[string]interface{}

func NewSitemapURL(url interface{}) sitemapURL {
	return sitemapURL{data: url.(URL)}
}

type sitemapURL struct {
	data URL
}

func (su sitemapURL) initialize() {
}

// craete  validators methods
// valid_keys

func (su sitemapURL) Xml() string {
	doc := etree.NewDocument()
	url := doc.CreateElement("url")
	priority := url.CreateElement("priority")
	priority.SetText(fmt.Sprint(4.2))
	_ = url.CreateElement("mobile:mobile")

	// if su.url.Priority > 0 {
	// priority := url.CreateElement("priority")
	// priority.SetText(fmt.Sprint("%f", su.url.Priority))
	// }

	// if su.url.Changefreq != "" {
	// changefreq := url.CreateElement("changefreq")
	// changefreq.SetText(su.url.Changefreq)
	// }

	// if su.url.Mobile {
	// _ = url.CreateElement("mobile:mobile")
	// }

	buf := &bytes.Buffer{}
	// doc.Indent(2)
	doc.WriteTo(buf)

	st := buf.String()

	// pretty.Println(st)
	// println("")

	return st
}
