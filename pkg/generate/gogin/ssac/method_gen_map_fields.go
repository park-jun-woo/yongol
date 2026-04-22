//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what methodGen.mapFields — Inputs map → "Key: value, ..." 문자열

package ssac

import (
	"sort"
	"strings"
)

// mapFields converts Inputs map → "Key: value, Key: value" string.
func (g *methodGen) mapFields(inputs map[string]string) string {
	keys := make([]string, 0, len(inputs))
	for k := range inputs {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var parts []string
	for _, k := range keys {
		parts = append(parts, k+": "+g.mapValue(inputs[k]))
	}
	return strings.Join(parts, ", ")
}
