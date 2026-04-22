//ff:type feature=rule type=model
//ff:what BaseSpec — 모든 Spec이 공유하는 Rule/Level/Message 필드
package rule

// BaseSpec provides common fields for all rule specs.
// Embed this in concrete spec types to satisfy toulmin.Spec.
type BaseSpec struct {
	Rule    string
	Level   string
	Message string
}
