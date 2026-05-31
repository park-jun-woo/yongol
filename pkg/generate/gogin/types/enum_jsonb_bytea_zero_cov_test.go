//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestHelpersZeroCov — head 판정 / native·pgtype·array·enum·jsonb·bytea binding 헬퍼 직접 커버
package types

import (
	"testing"
)

func TestEnumJsonbBytea_ZeroCov(t *testing.T) {
	if b := enumBinding(true, "'pending'"); b.Kind != KindEnum || b.ApiField != "string" {
		t.Errorf("enum NOT NULL = %+v", b)
	}
	if b := enumBinding(false, "'pending'"); b.ApiField != "*string" {
		t.Errorf("enum nullable = %+v", b)
	}
	if b := jsonbBinding(true, "'{}'"); b.Kind != KindJSONB || b.ApiField != "map[string]interface{}" {
		t.Errorf("jsonb NOT NULL = %+v", b)
	}
	if b := jsonbBinding(false, "'{}'"); b.ApiField != "*map[string]interface{}" {
		t.Errorf("jsonb nullable = %+v", b)
	}
	if b := byteaBinding(true, "''"); b.Kind != KindBytea || b.SqlcGoType != "[]byte" {
		t.Errorf("bytea = %+v", b)
	}
	if b := byteaBinding(false, "''"); b.Kind != KindBytea {
		t.Errorf("bytea nullable = %+v", b)
	}
}
