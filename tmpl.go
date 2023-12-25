package tmpl

import "strings"

type Tmpl[T any] struct {
	tmpl    string
	mapping func(T) map[string]string
}

func New[T any](tmpl string, mapping func(T) map[string]string) Tmpl[T] {
	return Tmpl[T]{
		tmpl:    tmpl,
		mapping: mapping,
	}
}

func (t Tmpl[T]) Render(params T) string {
	res := t.tmpl
	for k, v := range t.mapping(params) {
		res = strings.ReplaceAll(res, k, v)
	}

	return res
}
