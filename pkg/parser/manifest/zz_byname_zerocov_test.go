//ff:func feature=projectconfig type=test control=sequence
//ff:what TestByName_ZeroCov — manifest auth 헬퍼들을 이름으로 직접 호출해 커버리지 귀속
package manifest

import (
	"testing"
)

func TestByNameAuthLines_ZeroCov(t *testing.T) {
	data := []byte(`version: "1.0"
backend:
  auth:
    user_table: users
    claims:
      sub: {}
      role: {}
    roles:
      - admin
      - user
`)
	node := FindAuthNode(data)
	if node == nil {
		t.Fatalf("FindAuthNode returned nil")
	}

	claimLines := map[string]int{}
	collectClaimLines(node, claimLines)
	if len(claimLines) != 2 {
		t.Errorf("collectClaimLines = %v, want 2", claimLines)
	}

	roleLines := map[string]int{}
	collectRoleLines(node, roleLines)
	if len(roleLines) != 2 {
		t.Errorf("collectRoleLines = %v, want 2", roleLines)
	}

	cl, rl := extractAuthLines(data)
	if len(cl) != 2 || len(rl) != 2 {
		t.Errorf("extractAuthLines = %v / %v", cl, rl)
	}

	if line := extractUserTableLine(data); line == 0 {
		t.Errorf("extractUserTableLine = 0")
	}

	// nil/absent auth paths
	if FindAuthNode([]byte("version: \"1.0\"\n")) != nil {
		t.Errorf("FindAuthNode without backend should be nil")
	}
	emptyC, emptyR := extractAuthLines([]byte("foo: bar\n"))
	if len(emptyC) != 0 || len(emptyR) != 0 {
		t.Errorf("extractAuthLines without auth should be empty")
	}
	if extractUserTableLine([]byte("foo: bar\n")) != 0 {
		t.Errorf("extractUserTableLine without auth should be 0")
	}
}
