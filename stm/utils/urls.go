package utils

import (
	"net/url"
	"strings"
)

// too slowly
func URLJoin(src string, joins ...string) string {
	var u *url.URL
	lastnum := len(joins)
	base, _ := url.Parse(src)

	for i, j := range joins {
		if !strings.HasSuffix(j, "/") && lastnum > (i+1) {
			j = j + "/"
		}

		u, _ = url.Parse(j)
		base = base.ResolveReference(u)
	}

	return base.String()
}
