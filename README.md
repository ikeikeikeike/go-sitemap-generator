# go-sitemap-generator

```go
package main

import (
	"github.com/ikeikeikeike/go-sitemap-generator/sitemap"
	// "github.com/kr/pretty"
)

func main() {
	sm := sitemap.NewSitemap()
	sm.SetDefaultHost("http://myhost.com")
	sm.SetSitemapsPath("sitemap/myhost.com")

	// sm.Create().
		// Add(sitemap.AddFunc(func(ctx *sitemap.Context) {
		// ctx.Priority = nil
		// ctx.Changefreq = nil
		// ctx.Lastmod = nil
	// })).
		// Add(sitemap.LineFunc(func(ctx *sitemap.Context) {
		// ctx.Priority = 1
		// ctx.Changefreq = 1
		// ctx.Lastmod = ""
	// }))

	sm.Create().
		Add(sitemap.URL{Changefreq: "1", Mobile: true}).
		Add(sitemap.URL{Changefreq: "2", Mobile: true})

	// pretty.Println(sm)

	sm.PingSearchEngines()
}
```
