//ff:func feature=gen-react type=test control=sequence
//ff:what TestByName_ZeroCov — react/stml 코드젠 헬퍼들을 이름으로 직접 호출해 커버리지 귀속
package stml

import (
	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
)

func byNameConstraints() map[string]map[string]oapiparser.FieldConstraint {
	mn := 1
	mx := 50
	lo := 0.0
	hi := 100.0
	return map[string]map[string]oapiparser.FieldConstraint{
		"CreateItem": {
			"Name":  {Type: "string", MinLength: &mn, MaxLength: &mx, Required: true},
			"Count": {Type: "integer", Minimum: &lo, Maximum: &hi},
		},
		"Login": {
			"Email":    {Type: "string", Format: "email", Required: true},
			"Password": {Type: "string", MinLength: &mn},
		},
	}
}
