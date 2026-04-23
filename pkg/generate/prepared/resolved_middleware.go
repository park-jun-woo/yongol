//ff:type feature=generate type=model
//ff:what Middleware — 미들웨어 해석 결과 (placeholder, Stage 5 확장)

package prepared

// Middleware is the derived shape for a single middleware entry.
// Phase001 leaves this as a name-only placeholder; Stage 5 fills in the
// resolved ordering, OpenAPI securityScheme binding, and config per
// middleware. Emitters currently still read manifest.Backend.Middleware
// directly — migrations land incrementally.
type Middleware struct {
	Name string
}
