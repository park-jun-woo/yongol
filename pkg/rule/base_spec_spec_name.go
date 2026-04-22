//ff:func feature=rule type=accessor
//ff:what SpecName — 규칙 ID 반환 (toulmin.Spec 구현)
package rule

func (s BaseSpec) SpecName() string { return s.Rule }
