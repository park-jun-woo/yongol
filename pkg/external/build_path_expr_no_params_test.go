//ff:func feature=external type=test control=sequence
//ff:what TestMethodInfo* — pathParams/bodyParams/buildPathExpr/signature/implementation 검증
package external

import (
	"testing"
)

func TestBuildPathExprNoParams(t *testing.T) {
	m := methodInfo{Path: "/items"}
	if got := m.buildPathExpr(); got != `"/items"` {
		t.Errorf("buildPathExpr = %q, want \"/items\"", got)
	}
}
