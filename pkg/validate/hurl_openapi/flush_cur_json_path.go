//ff:func feature=validate type=util control=sequence topic=hurl-openapi
//ff:what flushCurJSONPath — cur 에 축적된 식별자를 out 으로 flush 하고 reset

package hurl_openapi

import "strings"

// flushCurJSONPath appends cur.String() to out when non-empty and
// resets the builder. Used by parseJSONPath / parseJSONPathChar when a
// separator (. or [) terminates the current identifier.
func flushCurJSONPath(cur *strings.Builder, out *[]string) {
	if cur.Len() == 0 {
		return
	}
	*out = append(*out, cur.String())
	cur.Reset()
}
