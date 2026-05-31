//ff:func feature=external type=test control=sequence
//ff:what TestExtractMethods/buildMethodInfo/extractResponseTypes — 메서드·응답타입 추출 검증
package external

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestExtractMethodsNilPaths(t *testing.T) {
	if got := extractMethods(&openapi3.T{}); len(got) != 0 {
		t.Errorf("expected no methods for nil Paths, got %+v", got)
	}
}
