//ff:func feature=gen-react type=test control=sequence
//ff:what detectDoublePrefix 일부 path 만 접두를 공유하면 이중 접두가 아니므로 nil (BUG-110)

package react

import "testing"

func TestDetectDoublePrefix_PartialPrefix_Nil(t *testing.T) {
	paths := []string{"/api/health", "/health"}
	if err := detectDoublePrefix("/api", paths); err != nil {
		t.Fatalf("expected nil when not every path shares the prefix, got %v", err)
	}
}
