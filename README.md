###### Inspired by [sitemap_generator](http://github.com/kjvarga/sitemap_generator)

##### How do I generate sitemap in Golang?

```go
package main

import (
    "github.com/ikeikeikeike/go-sitemap-generator/stm"
)


func main() {
    sm := stm.NewSitemap()
    sm.SetDefaultHost("http://example.com")
    sm.SetSitemapsPath("sitemap/example.com")

    sm.Create()

    sm.Add(stm.URL{"loc": "home", "changefreq": "always", "mobile": true})
    sm.Add(stm.URL{"loc": "readme"})
    sm.Add(stm.URL{"loc": "aboutme", "priority": 0.1})

    sm.Finalize().PingSearchEngines()
}
```

### Installing

```console
$ go get github.com/ikeikeikeike/go-sitemap-generator/stm
```

## Getting Started

### Preventing Output

```go
sm := stm.NewSitemap()
sm.SetVerbose(false)
```

### Pinging Search Engines

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

// Struct of `stm.S3Adapter`
sm.SetAdapter(stm.NewS3Adapter())

// It changes to output filename
sm.SetFilename("new_filename")
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
    stm.URL{"loc": "http://www.example.com/image.png", "title": "Image"},
    stm.URL{"loc": "http://www.example.com/image1.png", "title": "Image1"},
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

### PageMap Sitemaps

```go
```

### Alternate Links

```go
```

### Mobile Sitemaps

```go
sm.Add(stm.URL{"loc": "mobiles", "mobile": true})
```

### How to testing

Prepare testing

```console
$ go get github.com/clbanning/mxj
```

Do testing

```console
$ (cd ./stm ; go test -v github.com/ikeikeikeike/go-sitemap-generator/stm...)
```

#### Inspired by [sitemap_generator](http://github.com/kjvarga/sitemap_generator)
