//ff:type feature=gen-hurl type=model
//ff:what detectedAuthOp — 감지된 auth operation 메타데이터

package hurl

// detectedAuthOp captures the minimum metadata smoke.hurl needs to emit
// an auth step: operationId + HTTP method + URL path. method is always
// "POST" under yongol's DSL (signup/login never use GET).
type detectedAuthOp struct {
	OpID   string
	Path   string
	Method string
}
