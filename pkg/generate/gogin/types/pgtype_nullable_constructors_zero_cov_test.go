//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestPgtypeConstructorsZeroCov — 모든 pgtype 생성자 + unsupportedBinding 직접 커버
package types

import (
	"testing"
)

func TestPgtypeNullableConstructors_ZeroCov(t *testing.T) {
	// inet / interval / numeric / uuid each branch on notNull for the api field.
	if b := pgtypeInet(true, ""); b.ApiField != "string" {
		t.Errorf("pgtypeInet NOT NULL api = %q", b.ApiField)
	}
	if b := pgtypeInet(false, ""); b.ApiField != "*string" {
		t.Errorf("pgtypeInet nullable api = %q", b.ApiField)
	}
	if b := pgtypeInterval(true, ""); b.ApiField != "string" {
		t.Errorf("pgtypeInterval NOT NULL api = %q", b.ApiField)
	}
	if b := pgtypeInterval(false, ""); b.ApiField != "*string" {
		t.Errorf("pgtypeInterval nullable api = %q", b.ApiField)
	}
	if b := pgtypeNumeric(true, ""); b.ApiField != "string" {
		t.Errorf("pgtypeNumeric NOT NULL api = %q", b.ApiField)
	}
	if b := pgtypeNumeric(false, ""); b.ApiField != "*string" {
		t.Errorf("pgtypeNumeric nullable api = %q", b.ApiField)
	}
	if b := pgtypeUUID(true, ""); b.ApiField != "openapi_types.UUID" {
		t.Errorf("pgtypeUUID NOT NULL api = %q", b.ApiField)
	}
	if b := pgtypeUUID(false, ""); b.ApiField != "*openapi_types.UUID" {
		t.Errorf("pgtypeUUID nullable api = %q", b.ApiField)
	}
}
