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

#### How to testing

```
$ (cd ./stm ; go test -v github.com/ikeikeikeike/go-sitemap-generator/stm...)
```

#### Inspired by 

[sitemap_generator](http://github.com/kjvarga/sitemap_generator) 
