//ff:type feature=gen-gogin type=test-helper
//ff:what opSpec — 테스트에서 buildDoc이 받는 최소 operation 기술 구조체
package boot

import "github.com/getkin/kin-openapi/openapi3"

// opSpec is a compact description of one OpenAPI operation used by
// buildDoc to assemble minimal test fixtures.
type opSpec struct {
	path, method, opID string
	sec                *openapi3.SecurityRequirements
}
