//ff:type feature=tsx-parser type=model
//ff:what ArgBinding — apiClient.<op>({ ... }) 객체 인자의 key/value 한 쌍

package tsx

// ArgBinding is a single property key inside apiClient.<op>({ ... }).
// Value is the raw source snippet (best-effort) used purely for diagnostics;
// XOT-2 only compares Key names against OpenAPI parameter names.
type ArgBinding struct {
	Key   string
	Value string
}
