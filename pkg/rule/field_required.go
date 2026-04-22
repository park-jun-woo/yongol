//ff:func feature=rule type=rule control=sequence
//ff:what FieldRequired — 필드 존재/부재 제약 검증
package rule

import "github.com/park-jun-woo/toulmin/pkg/toulmin"

// FieldRequired checks that a field is present or absent.
// claim: map[string]bool (field name -> has value).
// Required=true: must be present. Required=false: must be absent.
func FieldRequired(ctx toulmin.Context, specs toulmin.Specs) (bool, any) {
	if len(specs) == 0 {
		return false, nil
	}
	s, ok := specs[0].(*FieldRequiredSpec)
	if !ok || s == nil {
		return false, nil
	}
	c, _ := ctx.Get("claim")
	fields, _ := c.(map[string]bool)
	present := fields[s.Field]
	if s.Required && !present {
		return true, &Evidence{Rule: s.Rule, Level: s.Level, Ref: s.Field, Message: s.Message}
	}
	if !s.Required && present {
		return true, &Evidence{Rule: s.Rule, Level: s.Level, Ref: s.Field, Message: s.Message}
	}
	return false, nil
}
