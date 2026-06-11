//ff:func feature=stml-gen type=test-helper control=sequence
//ff:what prefillConstraints — prefill 렌더 테스트용 requestBody 제약조건 fixture

package stml

import oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"

// prefillConstraints returns the UpdateRule field constraints used by the
// data-prefill render tests.
func prefillConstraints() map[string]map[string]oapiparser.FieldConstraint {
	maxLen := 200
	return map[string]map[string]oapiparser.FieldConstraint{
		"UpdateRule": {
			"sheet_name": {Type: "string", Required: true, MaxLength: &maxLen},
			"start_row":  {Type: "integer", Required: true},
			"note":       {Type: "string"},
		},
	}
}
