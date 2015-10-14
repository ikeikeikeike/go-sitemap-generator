# go-sitemap-generator
```go
package main

import (
	"github.com/ikeikeikeike/go-sitemap-generator/stm"
)

func main() {
	sm := stm.NewSitemap()
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

	for i := 0; i < 30000; i++ {
		builder.Add(stm.URL{"changefreq": "1", "mobile": true})
	}

	//pretty.Println(builder.Content())

	sm.PingSearchEngines(builder)
}
```
