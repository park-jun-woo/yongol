//ff:func feature=rule type=test control=sequence
//ff:what hasBearerAuthMiddleware — bearerAuth 포함 여부 pure 헬퍼 회귀

package ground

import (
	"testing"
)

// TestHasBearerAuthMiddleware covers the small pure helper.
func TestHasBearerAuthMiddleware(t *testing.T) {
	if !hasBearerAuthMiddleware([]string{"cors", "bearerAuth", "gzip"}) {
		t.Errorf("expected true when bearerAuth present")
	}
	if hasBearerAuthMiddleware([]string{"cors"}) {
		t.Errorf("expected false when absent")
	}
	if hasBearerAuthMiddleware(nil) {
		t.Errorf("expected false on nil")
	}
}
