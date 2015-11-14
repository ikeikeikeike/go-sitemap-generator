package stm

import (
	"testing"
	"time"

	"github.com/beevik/etree"
)

func TestBlank(t *testing.T) {
	if _, err := NewSitemapURL(URL{}); err == nil {
		t.Errorf(`Failed to validate blank arg ( URL{} ): %v`, err)
	}
}

func TestItHaveLocElement(t *testing.T) {
	if _, err := NewSitemapURL(URL{}); err == nil {
		t.Errorf(`Failed to validate about must have loc attribute in URL type ( URL{} ): %v`, err)
	}
}

func TestJustSetLocElement(t *testing.T) {
	smu, err := NewSitemapURL(URL{"loc": "path"})

	if err != nil {
		t.Fatalf(`Fatal to validate! This is a critical error: %v`, err)
	}

	doc := etree.NewDocument()
	doc.ReadFromBytes(smu.XML())

	var elm *etree.Element
	url := doc.SelectElement("url")

	elm = url.SelectElement("loc")
	if elm == nil {
		t.Errorf(`Failed to generate xml that loc attribute is blank: %v`, elm)
	}
	if elm != nil && elm.Text() != "path" {
		t.Errorf(`Failed to generate xml thats deferrent value in loc attribute: %v`, elm.Text())
	}
}

func TestJustSetLocElementAndThenItNeedsCompleteValues(t *testing.T) {
	smu, err := NewSitemapURL(URL{"loc": "path"})

	if err != nil {
		t.Fatalf(`Fatal to validate! This is a critical error: %v`, err)
	}

	doc := etree.NewDocument()
	doc.ReadFromBytes(smu.XML())

	var elm *etree.Element
	url := doc.SelectElement("url")

	elm = url.SelectElement("loc")
	if elm == nil {
		t.Errorf(`Failed to generate xml that loc element is blank: %v`, elm)
	}
	if elm != nil && elm.Text() != "path" {
		t.Errorf(`Failed to generate xml thats deferrent value in loc element: %v`, elm.Text())
	}

	elm = url.SelectElement("priority")
	if elm == nil {
		t.Errorf(`Failed to generate xml that priority element is nil: %v`, elm)
	}
	if elm != nil && elm.Text() != "0.5" {
		t.Errorf(`Failed to generate xml thats deferrent value in priority element: %v`, elm.Text())
	}

	elm = url.SelectElement("changefreq")
	if elm == nil {
		t.Errorf(`Failed to generate xml that changefreq element is nil: %v`, elm)
	}
	if elm != nil && elm.Text() != "weekly" {
		t.Errorf(`Failed to generate xml thats deferrent value in changefreq element: %v`, elm.Text())
	}

	elm = url.SelectElement("lastmod")
	if elm == nil {
		t.Errorf(`Failed to generate xml that lastmod element is nil: %v`, elm)
	}
	if elm != nil {
		if _, err := time.Parse(time.RFC3339, elm.Text()); err != nil {
			t.Errorf(`Failed to generate xml thats failed to parse datetime in lastmod element: %v`, err)
		}
	}
}

// func TestSetNilValue(t *testing.T) {
// smu, err := NewSitemapURL(URL{"loc": "path", "priority": nil})

// if err != nil {
// t.Fatalf(`Fatal to validate! This is a critical error: %v`, err)
// }

// doc := etree.NewDocument()
// doc.ReadFromBytes(smu.XML())

// var elm *etree.Element
// url := doc.SelectElement("url")

// elm = url.SelectElement("loc")
// if elm == nil {
// t.Errorf(`Failed to generate xml that loc attribute is blank: %v`, elm)
// }
// if elm != nil && elm.Text() != "path" {
// t.Errorf(`Failed to generate xml thats deferrent value in loc attribute: %v`, elm.Text())
// }

// pp.Println(smu.data)
// pp.Println(string(smu.XML()))
// }
