##### How I do generate sitemap in Golang?

```go
package main

import (
	"github.com/ikeikeikeike/go-sitemap-generator/stm"
)

func main() {
	sm := stm.NewSitemap()
	sm.SetDefaultHost("http://myhost.com")
	sm.SetSitemapsPath("sitemap/myhost.com")

	sm.Create()

	for i := 0; i < 30000; i++ {
		sm.Add(stm.URL{"changefreq": "1", "mobile": true})
	}

	sm.Finalize().PingSearchEngines()
}
```
