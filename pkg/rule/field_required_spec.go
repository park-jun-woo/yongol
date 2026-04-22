//ff:type feature=rule type=model
//ff:what FieldRequiredSpec — FieldRequired 규칙의 판정 기준
package rule

// FieldRequiredSpec configures a FieldRequired rule.
// Required=true means the field must be present; false means absent.
type FieldRequiredSpec struct {
	BaseSpec
	SeqType  string
	Field    string
	Required bool
}
