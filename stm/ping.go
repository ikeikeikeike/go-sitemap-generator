package stm

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func PingSearchEngines(bldr Builder, urls ...string) {
	urls = append(urls, []string{
		"http://www.google.com/webmasters/tools/ping?sitemap=%s",
		"http://www.bing.com/webmaster/ping.aspx?siteMap=%s",
		// "http://www.oogle.com/webmasters/tools/ping?sitemap=%s",
		// "http://www.ing.com/webmaster/ping.aspx?siteMap=%s",
		// "http://www.kdlakal.com/webmaster/ping.aspx?siteMap=%s",
	}...)

	nums := len(urls)
	does := make(chan string, nums)
	client := http.Client{Timeout: time.Duration(5 * time.Second)}

	for _, url := range urls {
		go func(url string) {
			log.Println("[I] Ping now:", url)

			resp, err := client.Get(url + "http://example.com/sitemap.tar.gz")
			if err != nil {
				does <- fmt.Sprintf("[E] Ping failed: %s (URL:%s)", err, url)
				return
			}
			defer resp.Body.Close()

			does <- fmt.Sprintf("[I] Successful ping of `%s`", url)
		}(url)
	}

	for i := 0; i < nums; i++ {
		log.Println(<-does)
	}
}
