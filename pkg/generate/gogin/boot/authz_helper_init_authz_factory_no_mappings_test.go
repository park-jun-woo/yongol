//ff:func feature=gen-gogin type=test control=sequence
//ff:what authzHelperInitAuthzFactory — OPA authz.Init 호출 헬퍼 (OwnershipMapping 임베드) 소스 생성
package boot

import (
	"strings"
	"testing"
)

func TestAuthzHelperInitAuthzFactory_NoMappings(t *testing.T) {
	src := authzHelperInitAuthzFactory(nil)
	if !strings.Contains(src, "authz.Init(policyPath, []authz.OwnershipMapping{") {
		t.Errorf("empty mappings should still emit Init call, got:\n%s", src)
	}
}
