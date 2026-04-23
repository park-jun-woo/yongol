//ff:type feature=generate type=model
//ff:what Route — 라우트 해석 결과 (placeholder, Stage 5 확장)

package prepared

// Route is the derived shape for a single route. Phase001 leaves this
// as a placeholder; Stage 5 fills in path, method, operationId, and the
// bound middleware list. Emitters currently still read raw OpenAPI /
// manifest data — migrations land incrementally.
type Route struct {
	OperationID string
	Method      string
	Path        string
}
