//ff:type feature=rule type=model
//ff:what Evidence — 규칙 위반 시 반환되는 결과 상세
package rule

// Evidence is the result detail returned by a violated rule.
type Evidence struct {
	Rule    string
	Level   string
	Ref     string
	Message string
}
