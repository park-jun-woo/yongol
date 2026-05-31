//ff:func feature=rule type=test control=sequence
//ff:what TestPopulateBatch_ZeroCov — ground populate/helper 함수를 이름으로 직접 호출해 커버 귀속
package ground

import (
	"testing"
)

func TestGroundPureHelpers_ZeroCov(t *testing.T) {
	if ctxIntDefault() != "int" {
		t.Error("ctxIntDefault")
	}
	if ctxNumberType("double") != "float64" || ctxNumberType("") != "float32" {
		t.Error("ctxNumberType")
	}
	if stringGoType("uuid", CtxResponseBody) != "openapi_types.UUID" {
		t.Error("stringGoType uuid")
	}
	if stringGoType("date-time", CtxResponseBody) != "time.Time" {
		t.Error("stringGoType date-time response")
	}
	if stringGoType("date-time", OAPIContext(99)) != "string" {
		t.Error("stringGoType date-time param")
	}
	if stringGoType("plain", CtxResponseBody) != "string" {
		t.Error("stringGoType default")
	}
}
