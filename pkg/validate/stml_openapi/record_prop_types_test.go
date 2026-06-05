//ff:func feature=validate type=test control=iteration dimension=1 topic=stml-openapi
//ff:what recordPropTypes — property→type 수집의 nil ref·nil value·미선언타입 폴백("") 분기 검증

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestRecordPropTypes(t *testing.T) {
	out := map[string]string{}
	props := openapi3.Schemas{
		"name":     typedSchema("string"),
		"age":      typedSchema("integer"),
		"untyped":  {Value: &openapi3.Schema{}}, // declared but no type → ""
		"nilRef":   nil,                         // nil ref → ""
		"nilValue": {Value: nil},                // nil value → ""
	}
	recordPropTypes(out, props)

	want := map[string]string{
		"name":     "string",
		"age":      "integer",
		"untyped":  "",
		"nilRef":   "",
		"nilValue": "",
	}
	for k, v := range want {
		if got, ok := out[k]; !ok || got != v {
			t.Errorf("out[%q] = %q (present=%v), want %q", k, got, ok, v)
		}
	}
}
