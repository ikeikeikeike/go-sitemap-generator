package stm

import (
	"fmt"
)

type BuilderError interface {
	error
	FullError() bool
}

type Builder interface {
	Content() []byte
	Add(interface{}) BuilderError
	Write()
	run()
}

type URL map[string]interface{}

func (u URL) URLJoinBy(key string, joins ...string) URL {
	var values []string
	for _, k := range joins {
		values = append(values, fmt.Sprint(u[k]))
	}

	u[key] = URLJoin("", values...)
	return u
}

func (u *URL) BungURLJoinBy(key string, joins ...string) {
	orig := *u

	var values []string
	for _, k := range joins {
		values = append(values, fmt.Sprint(orig[k]))
	}

	orig[key] = URLJoin("", values...)
	*u = orig
}

// type News map[string]interface{}
