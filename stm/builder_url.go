package stm

import (
	"bytes"
	"errors"
	"fmt"
	"time"

	"github.com/beevik/etree"
	"github.com/fatih/structs"
	"github.com/ikeikeikeike/go-sitemap-generator/stm/utils"
)

// http://www.sitemaps.org/protocol.html
// https://support.google.com/webmasters/answer/178636
type URLModel struct {
	Priority   float64                `valid:"float,length(0.0|1.0)"`
	Changefreq string                 `valid:"alpha(always|hourly|daily|weekly|monthly|yearly|never)"`
	Lastmod    time.Time              `valid:"-"`
	Expires    time.Time              `valid:"-"`
	Host       string                 `valid:"ipv4"`
	Loc        string                 `valid:"url"`
	Images     string                 `valid:"url"`
	Videos     string                 `valid:"url"`
	Geo        string                 `valid:"latitude,longitude"`
	News       string                 `valid:"-"`
	Mobile     bool                   `valid:"-"`
	Alternate  string                 `valid:"-"`
	Alternates map[string]interface{} `valid:"-"`
	Pagemap    map[string]interface{} `valid:"-"`
}

// []string{"priority" "changefreq" "lastmod" "expires" "host" "images"
// "video" "geo" "news" "videos" "mobile" "alternate" "alternates" "pagemap"}
var fieldnames []string = utils.ToLowers(structs.Names(&URLModel{}))

func NewSitemapURL(url interface{}) (*sitemapURL, error) {
	smu := &sitemapURL{data: url.(URL)}
	err := smu.validate()
	return smu, err
}

type sitemapURL struct {
	data URL
}

func (su *sitemapURL) validate() error {
	var key string
	var invalid bool

	for key, _ = range su.data {
		invalid = true
		for _, name := range fieldnames {
			if key == name {
				invalid = false
				break
			}
		}
		if invalid {
			break
		}
	}

	if invalid {
		msg := fmt.Sprintf("Unknown map key `%s` in URL type", key)
		return errors.New(msg)
	}
	if _, ok := su.data["loc"]; !ok {
		msg := fmt.Sprintf("URL type must have loc attribute")
		return errors.New(msg)
	}
	return nil
}

func (su *sitemapURL) XML() []byte {
	doc := etree.NewDocument()
	url := doc.CreateElement("url")

	utils.SetElementValue(url, su.data, "loc")
	utils.SetElementValue(url, su.data, "expires")
	utils.SetElementValue(url, su.data, "mobile")

	if !utils.SetElementValue(url, su.data, "changefreq") {
		changefreq := url.CreateElement("changefreq")
		changefreq.SetText("weekly")
	}
	if !utils.SetElementValue(url, su.data, "priority") {
		priority := url.CreateElement("priority")
		priority.SetText("0.5")
	}
	if !utils.SetElementValue(url, su.data, "lastmod") {
		lastmod := url.CreateElement("lastmod")
		lastmod.SetText(time.Now().Format(time.RFC3339))
	}

	buf := &bytes.Buffer{}
	// doc.Indent(2)
	doc.WriteTo(buf)

	return buf.Bytes()
}
