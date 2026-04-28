//ff:func feature=validate type=util control=sequence topic=ssac-structural
//ff:what splitEvalModel — "pkg.Method" → (pkg, Method) helper for @eval lookups

package ssac

import "strings"

// splitEvalModel splits an @eval Model string ("pkg.Method") into (pkg, Method).
// Returns ("", "") when the dot separator is absent.
func splitEvalModel(model string) (string, string) {
	idx := strings.IndexByte(model, '.')
	if idx <= 0 {
		return "", ""
	}
	return model[:idx], model[idx+1:]
}
