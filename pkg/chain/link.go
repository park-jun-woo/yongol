//ff:type feature=chain type=model
//ff:what Link — feature chain 의 단일 노드 (SSOT 또는 artifact)
package chain

// Link represents one SSOT or artifact node in a feature chain.
type Link struct {
	Kind      string // "OpenAPI", "SSaC", "DDL", "Rego", "StateDiag", "FuncSpec", "Hurl", "Handler", "Model", "Authz", "Types"
	File      string // specs-dir 기준 상대 경로
	Line      int    // 1-based line number, 0 if unknown
	Summary   string // brief description of the match
	Ownership string // "", "gen", "preserve" (empty for SSOT nodes)
}
