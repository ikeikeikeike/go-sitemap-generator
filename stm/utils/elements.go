package utils

import (
	"fmt"
	"time"

	"github.com/beevik/etree"
)

func SetElementValue(elm *etree.Element, data map[string]interface{}, key string) bool {
	if value, ok := data[key]; ok {
		switch v := value.(type) {
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
		default:
			child := elm.CreateElement(key)
			child.SetText(fmt.Sprint(v))
		case nil:
		}

		return true
	}
	return false
}
