package utils

import (
	"math/rand"

	"github.com/grafana/sobek"
	"go.k6.io/k6/js/modules"
)

var hexRunes = []rune("123456789abcdef")

func RandHexStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = hexRunes[rand.Intn(len(hexRunes))]
	}
	return string(b)
}

func ParseOptions(vu modules.VU, val sobek.Value) map[string]sobek.Value {
	options := make(map[string]sobek.Value)
	rt := vu.Runtime()

	if IsNilly(val) {
		return options
	}

	params := val.ToObject(rt)
	for _, k := range params.Keys() {
		options[k] = params.Get(k)
	}

	return options
}

func IsNilly(val sobek.Value) bool {
	return val == nil || sobek.IsNull(val) || sobek.IsUndefined(val)
}
