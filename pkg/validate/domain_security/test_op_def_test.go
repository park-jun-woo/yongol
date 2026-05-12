//ff:type feature=validate type=model
//ff:what opDef — 테스트용 OpenAPI operation 정의
package domain_security

// opDef defines an operation for test YAML generation.
type opDef struct {
	ID       string
	Security string // e.g. "[]" or "[{bearerAuth: []}]"
}
