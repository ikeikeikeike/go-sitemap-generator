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


#### How to testing

```
$ (cd ./stm ; go test -v github.com/ikeikeikeike/go-sitemap-generator/stm...)
```
