//ff:type feature=report type=model topic=sarif
//ff:what Region — 1-based line/column span (startColumn 은 0 이면 omit)
package sarif

// Region is a 1-based line/column span. startColumn omitted when 0.
type Region struct {
	StartLine   int `json:"startLine,omitempty"`
	StartColumn int `json:"startColumn,omitempty"`
}
