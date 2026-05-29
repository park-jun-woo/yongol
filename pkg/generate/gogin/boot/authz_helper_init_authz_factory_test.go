//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what authzHelperInitAuthzFactory — OPA authz.Init 호출 헬퍼 (OwnershipMapping 임베드) 소스 생성

package boot

import (
	"strings"
	"testing"
)

func TestAuthzHelperInitAuthzFactory(t *testing.T) {
	mappings := []string{
		`{Resource: "note", Table: "notes", Column: "owner_id"},`,
	}
	src := authzHelperInitAuthzFactory(mappings)
	for _, must := range []string{
		"func initAuthz(policyPath string) {",
		`if policyPath == "" {`,
		`slog.Error("OPA_POLICY_PATH is required")`,
		"os.Stat(policyPath)",
		"authz.Init(policyPath, []authz.OwnershipMapping{",
		`{Resource: "note", Table: "notes", Column: "owner_id"},`,
	} {
		if !strings.Contains(src, must) {
			t.Errorf("initAuthz factory missing %q, got:\n%s", must, src)
		}
	}
	// No DB threading post Phase002.
	if strings.Contains(src, "*sql.DB") || strings.Contains(src, "conn") {
		t.Errorf("initAuthz must be DB-free (Phase002), got:\n%s", src)
	}
}

func TestAuthzHelperInitAuthzFactory_NoMappings(t *testing.T) {
	src := authzHelperInitAuthzFactory(nil)
	if !strings.Contains(src, "authz.Init(policyPath, []authz.OwnershipMapping{") {
		t.Errorf("empty mappings should still emit Init call, got:\n%s", src)
	}
}
