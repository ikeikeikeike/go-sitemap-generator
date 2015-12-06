###### Inspired by [sitemap_generator](http://github.com/kjvarga/sitemap_generator)

##### How do I generate sitemap in Golang?  

[![GoDoc](https://godoc.org/github.com/ikeikeikeike/go-sitemap-generator/stm?status.svg)](https://godoc.org/github.com/ikeikeikeike/go-sitemap-generator/stm) [![Build Status](https://travis-ci.org/ikeikeikeike/go-sitemap-generator.svg)](https://travis-ci.org/ikeikeikeike/go-sitemap-generator)

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

```go
package main

import (
    "github.com/ikeikeikeike/go-sitemap-generator/stm"
)

func main() {
    sm := stm.NewSitemap()
    sm.SetDefaultHost("http://example.com")
    sm.SetSitemapsPath("sitemap-generator")  // default: public
    sm.SetSitemapsHost("http://s3.amazonaws.com/sitemap-generator/")
    sm.SetAdapter(&stm.S3Adapter{
        Region: "ap-northeast-1", 
        Bucket: "your-bucket", 
        ACL: "public-read",
    })

    sm.Create()

    sm.Add(stm.URL{"loc": "home", "changefreq": "always", "mobile": true})
    sm.Add(stm.URL{"loc": "readme"})
    sm.Add(stm.URL{"loc": "aboutme", "priority": 0.1})

    sm.Finalize().PingSearchEngines()
}
```

### News Sitemaps

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

### Image Sitemaps

```go
sm.Add(stm.URL{"loc": "/images", "image": []stm.URL{
    {"loc": "http://www.example.com/image.png", "title": "Image"},
    {"loc": "http://www.example.com/image1.png", "title": "Image1"},
}})

```

### Video Sitemaps

```go
sm.Add(stm.URL{"loc": "/videos", "video": stm.URL{
    "thumbnail_loc": "http://www.example.com/video1_thumbnail.png",
    "title":         "Title",
    "description":   "Description",
    "content_loc":   "http://www.example.com/cool_video.mpg",
    "category":      "Category",
    "tag":          []string{"one", "two", "three"},
}})
```

### Geo Sitemaps

```go
sm.Add(stm.URL{"loc": "/geos", "geo": stm.URL{
    "format": "kml",
}})
```

### Mobile Sitemaps

```go
sm.Add(stm.URL{"loc": "mobiles", "mobile": true})
```

### Example


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

	sm.Add(stm.URL{"loc": "/home", "changefreq": "dayly"})

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
	}})

	sm.Add(stm.URL{"loc": "/geos", "geo": stm.URL{
		"format": "kml",
	}})

	sm.Finalize().PingSearchEngines("http://newengine.com/ping?url=%s")
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
$ go test -v -cover ./... 
```

#### Inspired by [sitemap_generator](http://github.com/kjvarga/sitemap_generator)
