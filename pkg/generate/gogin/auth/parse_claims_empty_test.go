//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestParseClaims — claims 맵→정렬된 ClaimField + 기본 타입 검증
package auth

import (
	"testing"
)

func TestParseClaimsEmpty(t *testing.T) {
	if got := parseClaims(nil); len(got) != 0 {
		t.Errorf("expected empty slice for nil claims, got: %v", got)
	}
}
