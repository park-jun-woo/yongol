//ff:func feature=generate type=util control=iteration dimension=1
//ff:what middlewaresFor — manifest.backend.middleware 이름 목록 해석 (placeholder)

package prepared

import "github.com/park-jun-woo/yongol/pkg/yongol"

// middlewaresFor produces a name-only Middleware slice from
// manifest.backend.middleware. Phase001 keeps the shape trivial —
// Stage 5 extends it with ordering, OpenAPI securityScheme binding,
// and per-middleware config as emitters are migrated one by one.
func middlewaresFor(fs *yongol.Fullstack) []Middleware {
	if fs == nil || fs.Manifest == nil {
		return nil
	}
	names := fs.Manifest.Backend.Middleware
	out := make([]Middleware, 0, len(names))
	for _, name := range names {
		out = append(out, Middleware{Name: name})
	}
	return out
}
