
A go-sitemap-generator is the easiest way to generate Sitemaps in Go.

As of version 2.0.0, This Repo is available as a [Go module](https://github.com/golang/go/wiki/Modules).


[![GoDoc](https://godoc.org/github.com/ikeikeikeike/go-sitemap-generator/stm?status.svg)](https://godoc.org/github.com/ikeikeikeike/go-sitemap-generator/stm) [![Build Status](https://travis-ci.org/ikeikeikeike/go-sitemap-generator.svg)](https://travis-ci.org/ikeikeikeike/go-sitemap-generator)

```go
package main

import (
	"github.com/ikeikeikeike/go-sitemap-generator/v2/stm"
)


func main() {
	sm := stm.NewSitemap(1)

	// Create method must be called first before adding entries to
	// the sitemap.
	sm.Create()

	sm.Add(stm.URL{{"loc", "home"}, {"changefreq", "always"}, {"mobile", true}})
	sm.Add(stm.URL{{"loc", "readme"}})
	sm.Add(stm.URL{{"loc", "aboutme"}, {"priority", 0.1}})

	sm.Finalize().PingSearchEngines()
}
```

Then

```console
$ go build
```

#### Installation (Legacy download instead of a [Go module](https://github.com/golang/go/wiki/Modules).)

```console
$ go get gopkg.in/ikeikeikeike/go-sitemap-generator.v1/stm
$ go get gopkg.in/ikeikeikeike/go-sitemap-generator.v2/stm
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
  - [x] Alternate Links
- [ ] Supports: adapters for sitemap storage.
  - [x] Filesystem
  - [x] [S3](#upload-sitemap-to-s3)
- [x] [Customizable sitemap working](#preventing-output)
- [x] [Notifies search engines (Google, Bing) of new sitemaps](#pinging-search-engines)
- [x] [Gives you complete control over your sitemap contents and naming scheme](#full-example)


## Getting Started

### Setting concurrency
To disable concurrency, set number of CPUs to 1.
```go
sm := stm.NewSitemap(1)
```

If you want to set max CPUs that are available, set number of CPUs <= 0.
```go
sm := stm.NewSitemap(0)
```

### Preventing Output

To disable all non-essential output you can set `sm.SetVerbose` to `false`.
To disable output inline use the following:

```go
sm := stm.NewSitemap(1)
sm.SetVerbose(false)
```

### Pinging Search Engines

PingSearchEngines notifies search engines of changes once a sitemap
has been generated or changed. The library will append Google and Bing to any engines passed in to the function.

```go
sm.Finalize().PingSearchEngines()
```

If you want to add `new search engine`, you can pass that in to the function:

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

// Change the output filename
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
	sm := stm.NewSitemap(1)
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

	sm.Add(stm.URL{{"loc", "home"}, {"changefreq", "always"}, {"mobile", true}})
	sm.Add(stm.URL{{"loc", "readme"}})
	sm.Add(stm.URL{{"loc", "aboutme"}, {"priority", 0.1}})

	sm.Finalize().PingSearchEngines()
}
```

### News sitemaps

```go
sm.Add(stm.URL{
	{"loc", "/news"},
	{"news", stm.URL{
	{"publication", stm.URL{
		{"name",     "Example"},
		{"language", "en"},
	},
	},
	{"title",            "My Article"},
	{"keywords",         "my article, articles about myself"},
	{"stock_tickers",    "SAO:PETR3"},
	{"publication_date", "2011-08-22"},
	{"access",           "Subscription"},
	{"genres",           "PressRelease"},
},},})
```

Look at [Creating a Google News Sitemap](https://support.google.com/news/publisher/answer/74288) as required.

### Video sitemaps

```go
sm.Add(stm.URL{
	{"loc", "/videos"},
	{"video", stm.URL{
	{"thumbnail_loc", "http://www.example.com/video1_thumbnail.png"},
	{"title",         "Title"},
	{"description",   "Description"},
	{"content_loc",   "http://www.example.com/cool_video.mpg"},
	{"category",      "Category"},
	{"tag",           []string{"one", "two", "three"}},
    {"player_loc",    stm.Attrs{"https://example.com/p/flash/moogaloop/6.2.9/moogaloop.swf?clip_id=26", map[string]string{"allow_embed": "Yes", "autoplay": "autoplay=1"}},},
},
},
})
```

Look at [Video sitemaps](https://support.google.com/webmasters/answer/80471) as required.

### Image sitemaps

```go
sm.Add(stm.URL{
	{"loc", "/images"},
	{"image", []stm.URL{
	{{"loc", "http://www.example.com/image.png"}, {"title", "Image"}},
	{{"loc", "http://www.example.com/image1.png"}, {"title", "Image1"}},
},},
})

```

Look at [Image sitemaps](https://support.google.com/webmasters/answer/178636) as required.

### Geo sitemaps

```go
sm.Add(stm.URL{
	{"loc", "/geos"},
	{"geo", stm.URL{
	{"format", "kml"},
},},
})
```

Couldn't find Geo sitemaps example, although it's similar to:

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
sm.Add(stm.URL{{"loc", "mobiles"}, {"mobile", true}})
```

Look at [Feature phone sitemaps](https://support.google.com/webmasters/answer/6082207) as required.


### Full example

```go
package main

import (
	"github.com/ikeikeikeike/go-sitemap-generator/stm"
)

func main() {
	sm := stm.NewSitemap(0)
	sm.SetDefaultHost("http://yourhost.com")
	sm.SetSitemapsHost("http://s3.amazonaws.com/sitemaps/")
	sm.SetSitemapsPath("sitemaps/")
	sm.SetFilename("anothername")
	sm.SetCompress(true)
	sm.SetVerbose(true)
	sm.SetAdapter(&stm.S3Adapter{Region: "ap-northeast-1", Bucket: "your-bucket"})

	sm.Create()

	sm.Add(stm.URL{{"loc", "/home"}, {"changefreq", "daily"}})

	sm.Add(stm.URL{{"loc", "/abouts"}, {"mobile", true}})

	sm.Add(stm.URL{{"loc", "/news"},
	{"news", stm.URL{
		{"publication", stm.URL{
			{"name",     "Example"},
			{"language", "en"},
		},
		},
		{"title",            "My Article"},
		{"keywords",         "my article, articles about myself"},
		{"stock_tickers",    "SAO:PETR3"},
		{"publication_date", "2011-08-22"},
		{"access",           "Subscription"},
		{"genres",           "PressRelease"},
	},},
	})

	sm.Add(stm.URL{{"loc", "/images"},
	{"image", []stm.URL{
		{{"loc", "http://www.example.com/image.png"}, {"title", "Image"}},
		{{"loc", "http://www.example.com/image1.png"}, {"title", "Image1"}},
	},},
	})

	sm.Add(stm.URL{{"loc", "/videos"},
	{"video", stm.URL{
		{"thumbnail_loc", "http://www.example.com/video1_thumbnail.png"},
		{"title",         "Title"},
		{"description",   "Description"},
		{"content_loc",   "http://www.example.com/cool_video.mpg"},
		{"category",      "Category"},
		{"tag",           []string{"one", "two", "three"}},
	    {"player_loc",    stm.Attrs{"https://example.com/p/flash/moogaloop/6.2.9/moogaloop.swf?clip_id=26", map[string]string{"allow_embed": "Yes", "autoplay": "autoplay=1"}}},
	},},
	})

	sm.Add(stm.URL{{"loc", "/geos"},
	{"geo", stm.URL{
		{"format", "kml"},
	},},
	})

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
	sm := stm.NewSitemap(1)
	sm.SetDefaultHost("http://example.com")

	sm.Create()
	sm.Add(stm.URL{{"loc", "/"}, {"changefreq", "daily"}})

	// Note: Do not call `sm.Finalize()` because it flushes
	// the underlying data structure from memory to disk.

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

### How to test.

Preparation:

```console
$ go get github.com/clbanning/mxj
```

Run tests:

```console
$ go test -v -cover -race ./...
```

#### Inspired by [sitemap_generator](http://github.com/kjvarga/sitemap_generator)
