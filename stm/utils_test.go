package stm

import (
	"reflect"
	"testing"
)

func TestMergeMap(t *testing.T) {
	var src, dst, expect map[string]interface{}
	src = map[string]interface{}{"loc": "1", "changefreq": "2", "mobile": true, "host": "http://google.com"}
	dst = map[string]interface{}{"host": "http://example.com"}
	expect = map[string]interface{}{"loc": "1", "changefreq": "2", "mobile": true, "host": "http://google.com"}

	src = MergeMap(src, dst)

	if !reflect.DeepEqual(src, expect) {
		t.Fatalf("Failed to maps merge: deferrent map %v and %v", src, expect)
	}
}
