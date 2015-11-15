package stm

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/beevik/etree"
	"github.com/imdario/mergo"
)

// Change to News type def
func SetBuilderElementValue(elm *etree.Element, data map[string]interface{}, basekey string) bool {
	key := basekey
	ts, tk := spaceDecompose(elm.Tag)
	_, sk := spaceDecompose(elm.Space)

	if elm.Tag != "" && ts != "" && tk != "" {
		key = fmt.Sprintf("%s:%s", elm.Space, basekey)
	} else if sk != "" {
		key = fmt.Sprintf("%s:%s", sk, basekey)
	}

	if value, ok := data[basekey]; ok {
		switch v := value.(type) {
		case nil:
		default:
			child := elm.CreateElement(key)
			child.SetText(fmt.Sprint(v))
		case int:
			child := elm.CreateElement(key)
			child.SetText(fmt.Sprint(v))
		case string:
			child := elm.CreateElement(key)
			child.SetText(v)
		case float64, float32:
			child := elm.CreateElement(key)
			child.SetText(fmt.Sprint(v))
		case time.Time:
			child := elm.CreateElement(key)
			child.SetText(v.Format(time.RFC3339))
		case bool:
			_ = elm.CreateElement(fmt.Sprintf("%s:%s", key, key))
		case interface{}:
			switch mvalue := value.(type) {
			case URL:
				var childkey string
				if sk == "" {
					childkey = fmt.Sprintf("%s:%s", key, key)
				} else {
					childkey = fmt.Sprint(key)
				}

				child := elm.CreateElement(childkey)
				for mk, _ := range mvalue {
					SetBuilderElementValue(child, mvalue, mk)
				}
			}
		}

		return true
	}
	return false
}

// TODO: Slow function: It wants to change fast function
func MergeMap(src, dst map[string]interface{}) map[string]interface{} {
	mergo.MapWithOverwrite(&dst, src)
	return dst
}

func ToLowerString(befores []string) (afters []string) {
	for _, name := range befores {
		afters = append(afters, strings.ToLower(name))
	}
	return afters
}

// TODO: Too slowly
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


func spaceDecompose(str string) (space, key string) {
	colon := strings.IndexByte(str, ':')
	if colon == -1 {
		return "", str
	}
	return str[:colon], str[colon+1:]
}
