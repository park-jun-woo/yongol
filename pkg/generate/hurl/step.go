//ff:type feature=gen-hurl type=model
//ff:what step — hurl 스모크 테스트의 단일 요청 단계
package hurl

// step represents a single hurl request step.
type step struct {
	Comment     string // section comment (e.g. "# ===== Auth =====")
	Method      string
	Path        string
	QueryParams string // "limit=2&sort=id&direction=asc" (without leading '?')
	OperationID string
	NeedsAuth   bool
	TokenVar    string // hurl variable name for Authorization (e.g. "token", "token_admin")
	HasBody     bool
	BodyJSON    string
	StatusCode  int
	Captures    []capture
	Assertions  []string
}
