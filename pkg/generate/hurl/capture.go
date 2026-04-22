//ff:type feature=gen-hurl type=model
//ff:what capture — hurl [Captures] 엔트리 (변수명 + jsonpath)
package hurl

// capture holds a single [Captures] entry.
type capture struct {
	VarName  string // e.g. "workflow_id"
	JSONPath string // e.g. "$.workflow.id"
}
