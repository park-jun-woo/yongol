//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what goStringListMap — map[string][]string 를 Go 리터럴로 렌더 (deterministic)

package boot

import (
	"fmt"
	"sort"
	"strings"
)

// goStringListMap renders a map[string][]string as a Go literal. Keys are
// emitted in sorted order for deterministic codegen output. An empty or nil
// map yields `map[string][]string{}` so downstream struct assignment stays
// assignment-safe (no nil dereference when length-checking later).
func goStringListMap(m map[string][]string) string {
	if len(m) == 0 {
		return `map[string][]string{}`
	}
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var b strings.Builder
	b.WriteString("map[string][]string{\n")
	for _, k := range keys {
		b.WriteString(fmt.Sprintf("\t\t\t%q: %s,\n", k, goStringSlice(m[k])))
	}
	b.WriteString("\t\t}")
	return b.String()
}
