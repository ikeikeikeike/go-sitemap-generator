package stm

import (
	"reflect"
	"testing"
)

func TestMergeMap(t *testing.T) {
	var src, dst, expect [][]interface{}
	src = [][]interface{}{{"loc", "1"}, {"changefreq", "2"}, {"mobile", true}, {"host", "http://google.com"}}
	dst = [][]interface{}{{"host", "http://example.com"}}
	expect = [][]interface{}{{"loc", "1"}, {"changefreq", "2"}, {"mobile", true}, {"host", "http://google.com"}}

	src = MergeMap(src, dst)

	if !reflect.DeepEqual(src, expect) {
		t.Fatalf("Failed to maps merge: deferrent map \n%#v\n and \n%#v\n", src, expect)
	}
}
