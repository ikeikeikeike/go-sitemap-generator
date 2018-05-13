##### go-sitemap-generator is the easiest way to generate Sitemaps in Go.

[![GoDoc](https://godoc.org/github.com/ikeikeikeike/go-sitemap-generator/stm?status.svg)](https://godoc.org/github.com/ikeikeikeike/go-sitemap-generator/stm) [![Build Status](https://travis-ci.org/ikeikeikeike/go-sitemap-generator.svg)](https://travis-ci.org/ikeikeikeike/go-sitemap-generator)

#### Note: feature/concurrency blanch is experimental.

```go
package main

import (
	"github.com/ikeikeikeike/go-sitemap-generator/stm"
)


func main() {
	sm := stm.NewSitemap()

	// Create method must be that calls first this method in that before
	// call to Add method on this struct.
	sm.Create()

	sm.Add(stm.URL{"loc": "home", "changefreq": "always", "mobile": true})
	sm.Add(stm.URL{"loc": "readme"})
	sm.Add(stm.URL{"loc": "aboutme", "priority": 0.1})

	sm.Finalize().PingSearchEngines()
}
```

Sitemap provides interface for create sitemap xml file and that has convenient interface.
And also needs to use first Sitemap struct if it wants to use this package.


### Installing

```console
$ go get github.com/ikeikeikeike/go-sitemap-generator/stm
```

### Features

Current Features or To-Do

- [ ] Supports: generate kind of some sitemaps.
  - [x] [News sitemaps](#news-sitemaps)
  - [x] [Video sitemaps](#video-sitemaps)
  - [x] [Image sitemaps](#image-sitemaps)
  - [x] [Geo sitemaps](#geo-sitemaps)
  - [x] [Mobile sitemaps](#mobile-sitemaps)
  - [ ] PageMap sitemap
  - [ ] Alternate Links
- [ ] Supports: write some kind of filesystem and object storage.
  - [x] Filesystem
  - [x] [S3](#upload-sitemap-to-s3)
  - [ ]  Some adapter
- [x] [Customizable sitemap working](#preventing-output)
- [x] [Notifies search engines (Google, Bing) of new sitemaps](#pinging-search-engines)
- [x] [Gives you complete control over your sitemap contents and naming scheme](#full-example)


## Getting Started

### Preventing Output

To disable all non-essential output you can give `false` to `sm.SetVerbose`.
To disable output in-code use the following:

```go
sm := stm.NewSitemap()
sm.SetVerbose(false)
```

### Pinging Search Engines

PingSearchEngines requests some ping server.

```go
sm.Finalize().PingSearchEngines()
```

If you want to add `new search engine`, you can set that to method's arguments. like this.

```go
sm.Finalize().PingSearchEngines("http://newengine.com/ping?url=%s")
```

### Options

```go
// Your website's host name
sm.SetDefaultHost("http://www.example.com")

// The remote host where your sitemaps will be hosted
sm.SetSitemapsHost("http://s3.amazonaws.com/sitemap-generator/")

// The directory to write sitemaps to locally
sm.SetPublicPath("tmp/")

// Set this to a directory/path if you don't want to upload to the root of your `SitemapsHost`
sm.SetSitemapsPath("sitemaps/")

// Struct of `S3Adapter`
sm.SetAdapter(&stm.S3Adapter{Region: "ap-northeast-1", Bucket: "your-bucket", ACL: "public-read"})

// It changes to output filename
sm.SetFilename("new_filename")
```

### Upload sitemap to S3

Recently I disabled this module [here](https://github.com/ikeikeikeike/go-sitemap-generator/blob/master/stm/_adapter_s3.go).

```go
package main

import (
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/ikeikeikeike/go-sitemap-generator/stm"
)

func main() {
	sm := stm.NewSitemap()
	sm.SetDefaultHost("http://example.com")
	sm.SetSitemapsPath("sitemap-generator") // default: public
	sm.SetSitemapsHost("http://s3.amazonaws.com/sitemap-generator/")
	sm.SetAdapter(&stm.S3Adapter{
		Region: "ap-northeast-1",
		Bucket: "your-bucket",
		ACL:    "public-read",
		Creds:  credentials.NewEnvCredentials(),
	})

	sm.Create()

	sm.Add(stm.URL{"loc": "home", "changefreq": "always", "mobile": true})
	sm.Add(stm.URL{"loc": "readme"})
	sm.Add(stm.URL{"loc": "aboutme", "priority": 0.1})

	sm.Finalize().PingSearchEngines()
}
```

### News sitemaps

```go
sm.Add(stm.URL{"loc": "/news", "news": stm.URL{
	"publication": stm.URL{
		"name":     "Example",
		"language": "en",
	},
	"title":            "My Article",
	"keywords":         "my article, articles about myself",
	"stock_tickers":    "SAO:PETR3",
	"publication_date": "2011-08-22",
	"access":           "Subscription",
	"genres":           "PressRelease",
}})
```

Look at [Creating a Google News Sitemap](https://support.google.com/news/publisher/answer/74288) as required.

### Video sitemaps

```go
sm.Add(stm.URL{"loc": "/videos", "video": stm.URL{
	"thumbnail_loc": "http://www.example.com/video1_thumbnail.png",
	"title":         "Title",
	"description":   "Description",
	"content_loc":   "http://www.example.com/cool_video.mpg",
	"category":      "Category",
	"tag":           []string{"one", "two", "three"},
    "player_loc":    stm.Attrs{"https://example.com/p/flash/moogaloop/6.2.9/moogaloop.swf?clip_id=26", map[string]string{"allow_embed": "Yes", "autoplay": "autoplay=1"}},
}})
```

Look at [Video sitemaps](https://support.google.com/webmasters/answer/80471) as required.

### Image sitemaps

```go
sm.Add(stm.URL{"loc": "/images", "image": []stm.URL{
	{"loc": "http://www.example.com/image.png", "title": "Image"},
	{"loc": "http://www.example.com/image1.png", "title": "Image1"},
}})

```

Look at [Image sitemaps](https://support.google.com/webmasters/answer/178636) as required.

### Geo sitemaps

```go
sm.Add(stm.URL{"loc": "/geos", "geo": stm.URL{
	"format": "kml",
}})
```

Couldn't find Geo sitemaps example. Although its like a below.

```xml
<url>
	<loc>/geos</loc>
	<geo:geo>
		<geo:format>kml</geo:format>
	</geo:geo>
</url>
```

### Mobile sitemaps

```go
sm.Add(stm.URL{"loc": "mobiles", "mobile": true})
```

Look at [Feature phone sitemaps](https://support.google.com/webmasters/answer/6082207) as required.


### Full example

```go
package main

import (
	"github.com/ikeikeikeike/go-sitemap-generator/stm"
)

func main() {
	sm := stm.NewSitemap()
	sm.SetDefaultHost("http://yourhost.com")
	sm.SetSitemapsHost("http://s3.amazonaws.com/sitemaps/")
	sm.SetSitemapsPath("sitemaps/")
	sm.SetFilename("anothername")
	sm.SetCompress(true)
	sm.SetVerbose(true)
	sm.SetAdapter(&stm.S3Adapter{Region: "ap-northeast-1", Bucket: "your-bucket"})

	sm.Create()

	sm.Add(stm.URL{"loc": "/home", "changefreq": "daily"})

	sm.Add(stm.URL{"loc": "/abouts", "mobile": true})

	sm.Add(stm.URL{"loc": "/news", "news": stm.URL{
		"publication": stm.URL{
			"name":     "Example",
			"language": "en",
		},
		"title":            "My Article",
		"keywords":         "my article, articles about myself",
		"stock_tickers":    "SAO:PETR3",
		"publication_date": "2011-08-22",
		"access":           "Subscription",
		"genres":           "PressRelease",
	}})

	sm.Add(stm.URL{"loc": "/images", "image": []stm.URL{
		{"loc": "http://www.example.com/image.png", "title": "Image"},
		{"loc": "http://www.example.com/image1.png", "title": "Image1"},
	}})

	sm.Add(stm.URL{"loc": "/videos", "video": stm.URL{
		"thumbnail_loc": "http://www.example.com/video1_thumbnail.png",
		"title":         "Title",
		"description":   "Description",
		"content_loc":   "http://www.example.com/cool_video.mpg",
		"category":      "Category",
		"tag":           []string{"one", "two", "three"},
	    "player_loc":    stm.Attrs{"https://example.com/p/flash/moogaloop/6.2.9/moogaloop.swf?clip_id=26", map[string]string{"allow_embed": "Yes", "autoplay": "autoplay=1"}},
	}})

	sm.Add(stm.URL{"loc": "/geos", "geo": stm.URL{
		"format": "kml",
	}})

	sm.Finalize().PingSearchEngines("http://newengine.com/ping?url=%s")
}
```

### Webserver example


```go
package main

import (
	"log"
	"net/http"

	"github.com/ikeikeikeike/go-sitemap-generator/stm"
)

func buildSitemap() *stm.Sitemap {
	sm := stm.NewSitemap()
	sm.SetDefaultHost("http://example.com")

	sm.Create()
	sm.Add(stm.URL{"loc": "/", "changefreq": "daily"})

	// Note: Do not call `sm.Finalize()` because it flushes
	// the underlying datastructure from memory to disk.

	return sm
}

func main() {
	sm := buildSitemap()

	mux := http.NewServeMux()
	mux.HandleFunc("/sitemap.xml", func(w http.ResponseWriter, r *http.Request) {
		// Go's webserver automatically sets the correct `Content-Type` header.
		w.Write(sm.XMLContent())
		return
	})

	log.Fatal(http.ListenAndServe(":8080", mux))
}
```


### Documentation

- [API Reference](https://godoc.org/github.com/ikeikeikeike/go-sitemap-generator/stm)
- [sitemap_generator](http://github.com/kjvarga/sitemap_generator)

### How to testing

Prepare testing

```console
$ go get github.com/clbanning/mxj
```

Do testing

```console
$ go test -v -cover -race ./...
```

#### Inspired by [sitemap_generator](http://github.com/kjvarga/sitemap_generator)
