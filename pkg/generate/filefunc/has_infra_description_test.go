//ff:func feature=gen-filefunc type=test control=sequence
//ff:what TestHasInfraDescription — infra 설명 테이블 존재 여부 검증

package filefunc

import "testing"

func TestHasInfraDescription(t *testing.T) {
	if !hasInfraDescription("auth") {
		t.Errorf("expected 'auth' to have infra description")
	}
	if hasInfraDescription("nonexistent-package") {
		t.Errorf("expected unknown package to have no infra description")
	}
}
