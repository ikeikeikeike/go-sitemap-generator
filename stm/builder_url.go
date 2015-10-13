package stm

import (
	"bytes"
	"time"

	"github.com/beevik/etree"
	"github.com/kr/pretty"

    // "gopkg.in/go-playground/validator.v8"
)

type URLModel struct {
	Priority   float32   `validate:"required"`
	Changefreq string    `validate:"required"`
	Lastmod    time.Time `validate:"required"`
	Expires    time.Time `validate:"required"`
	Host       string    `validate:"required"`
	Loc        string    `validate:"required"`
	Images     string    `validate:"required"`
	Geo        string    `validate:"required"`
	Mobile     bool      `validate:"required"`
	Alternates string    `validate:"required"`
	Pagemap    string    `validate:"required"`
}

type URL map[string]interface{}

func NewSitemapURL(url interface{}) sitemapURL {
	pretty.Println(url)
	// u := url.(URL)
	// pretty.Println(structs.Map(u))
	// su := sitemapURL{url: url}
	// return su
	return sitemapURL{}
}

type sitemapURL struct {
	url URL
}

func (su sitemapURL) Xml() string {

	// mxj.Map()

	doc := etree.NewDocument()
	// url := doc.CreateElement("url")

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
	doc.Indent(2)
	doc.WriteTo(buf)

	st := buf.String()

	// pretty.Println(st)
	// println("")

	return st
}
