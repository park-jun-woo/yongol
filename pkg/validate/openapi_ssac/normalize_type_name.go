//ff:func feature=validate type=util control=sequence topic=openapi-ssac
//ff:what normalizeTypeName — slice/pointer/wrapper/pkg prefix 제거하여 canonical 타입명 반환

package openapi_ssac

import "strings"

// normalizeTypeName strips []/*, Wrapper[T] syntax, and package prefix to
// produce a canonical unqualified type name usable as Struct.<X>.<field> key.
//
//	"[]billing.CheckCreditsResponse" → "CheckCreditsResponse"
//	"*User"                          → "User"
//	"Page[Workflow]"                 → "Workflow"
func normalizeTypeName(t string) string {
	t = strings.TrimPrefix(t, "[]")
	t = strings.TrimPrefix(t, "*")
	// wrapper: Page[T], Cursor[T] → T
	if open := strings.IndexByte(t, '['); open > 0 && strings.HasSuffix(t, "]") {
		t = t[open+1 : len(t)-1]
	}
	if dot := strings.LastIndex(t, "."); dot >= 0 {
		t = t[dot+1:]
	}
	return t
}
