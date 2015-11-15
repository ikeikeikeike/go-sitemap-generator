package stm

import (
	"reflect"
	"testing"
)

func TestURLType(t *testing.T) {
	url := URL{"loc": "1", "host": "http://example.com"}
	expect := URL{"loc": "http://example.com/1", "host": "http://example.com"}

	url = url.URLJoinBy("loc", "host", "loc")

	if !reflect.DeepEqual(url, expect) {
		t.Fatalf("Failed to join url in URL type: deferrent URL %v and %v", url, expect)
	}

	url = URL{"loc": "1", "host": "http://example.com", "mobile": true}
	expect = URL{"loc": "http://example.com/1/true", "host": "http://example.com", "mobile": true}

	url.BungURLJoinBy("loc", "host", "loc", "mobile")

	if !reflect.DeepEqual(url, expect) {
		t.Fatalf("Failed to join url in URL type: deferrent URL %v and %v", url, expect)
	}
}
