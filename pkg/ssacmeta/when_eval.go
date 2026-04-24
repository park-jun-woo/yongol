//ff:func feature=ssacmeta type=util control=selection
//ff:what EvaluateWhen — interface.yaml 의 when: 표현식을 manifest 컨텍스트로 평가

package ssacmeta

import "strings"

// EvaluateWhen returns true when the port should be active under the given
// manifest context. Supported syntax (minimal):
//
//	always                              — always active
//	manifest.<path> == "<value>"        — equal comparison on a manifest field
//	manifest.<path>                     — truthy check on a manifest field
//
// Paths are dot-separated keys into the flattened manifest map. Unknown
// paths evaluate to false.
func EvaluateWhen(expr string, manifest map[string]any) bool {
	e := strings.TrimSpace(expr)
	if e == "" || e == "always" {
		return true
	}
	if !strings.HasPrefix(e, "manifest.") {
		// Unknown prefix — conservative default: inactive.
		return false
	}
	// equality form: manifest.x.y == "value"
	if i := strings.Index(e, "=="); i > 0 {
		lhs := strings.TrimSpace(e[:i])
		rhs := strings.TrimSpace(e[i+2:])
		rhs = strings.Trim(rhs, `"`)
		v, ok := lookupPath(manifest, strings.TrimPrefix(lhs, "manifest."))
		if !ok {
			return false
		}
		return fmtLiteral(v) == rhs
	}
	// truthy form: manifest.x.y
	v, ok := lookupPath(manifest, strings.TrimPrefix(e, "manifest."))
	if !ok {
		return false
	}
	return truthy(v)
}

func lookupPath(m map[string]any, path string) (any, bool) {
	parts := strings.Split(path, ".")
	var cur any = m
	for _, p := range parts {
		mm, ok := cur.(map[string]any)
		if !ok {
			return nil, false
		}
		cur, ok = mm[p]
		if !ok {
			return nil, false
		}
	}
	return cur, true
}

func truthy(v any) bool {
	switch x := v.(type) {
	case nil:
		return false
	case bool:
		return x
	case string:
		return x != ""
	case int:
		return x != 0
	case int64:
		return x != 0
	case float64:
		return x != 0
	default:
		return true
	}
}

func fmtLiteral(v any) string {
	if s, ok := v.(string); ok {
		return s
	}
	return ""
}
