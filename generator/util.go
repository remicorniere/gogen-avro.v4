package generator

import (
	"sort"
	"strings"
	"unicode"

	"github.com/serenize/snaker"
)

// ToPublicName returns a go-idiomatic public name
func ToPublicName(name string) string {
	return snaker.SnakeToCamel(name)
}

func concatSortedMap(m map[string]string, sep string) string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for n := 0; n < len(m); n++ {
		keys[n] = m[keys[n]]
	}

	return strings.Join(keys, sep) + sep
}

// Make filenames snake-case, taken from https://gist.github.com/elwinar/14e1e897fdbe4d3432e1
func ToSnake(in string) string {
	runes := []rune(in)
	length := len(runes)

	var out []rune
	for i := 0; i < length; i++ {
		if i > 0 && unicode.IsUpper(runes[i]) && ((i+1 < length && unicode.IsLower(runes[i+1])) || unicode.IsLower(runes[i-1])) {
			out = append(out, '_')
		}
		out = append(out, unicode.ToLower(runes[i]))
	}

	return string(out)
}
