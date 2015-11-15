package utils

import "github.com/imdario/mergo"

// TODO: Slow function: It wants to change fast function
func MergeMap(src, dst map[string]interface{}) map[string]interface{} {
	mergo.MapWithOverwrite(&dst, src)
	return dst
}
