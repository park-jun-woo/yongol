//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestFastapiHelpers — fastapi plan/package/route 헬퍼 검증 (Op 종류·외부 패키지 수집·라우트 해석)
package fastapi

import (
	"testing"
)

func TestEnsurePkgMap(t *testing.T) {
	pm := map[string]map[string]bool{}
	ensurePkgMap(pm, "billing")
	if pm["billing"] == nil {
		t.Fatal("expected sub-map created")
	}
	pm["billing"]["X"] = true
	ensurePkgMap(pm, "billing") // idempotent
	if !pm["billing"]["X"] {
		t.Error("existing sub-map should be preserved")
	}
}
