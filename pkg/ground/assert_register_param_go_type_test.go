//ff:func feature=rule type=test-helper control=sequence
//ff:what assertRegisterParamGoType — registerParamGoType 의 OpenAPI.paramType 키 등록 여부/값 검증 헬퍼
package ground

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

// assertRegisterParamGoType registers param under op and asserts whether the
// OpenAPI.paramType key was registered and its value.
func assertRegisterParamGoType(t *testing.T, op, caseName string, param *openapi3.Parameter, wantType string, wantFound bool) {
	t.Helper()
	g := newGround()
	registerParamGoType(g, op, param)
	got, found := g.Types["OpenAPI.paramType."+op+"."+param.Name]
	if found != wantFound {
		t.Fatalf("%s: key present=%v, want %v (value=%q)", caseName, found, wantFound, got)
	}
	if found && got != wantType {
		t.Errorf("%s: registered %q, want %q", caseName, got, wantType)
	}
}
