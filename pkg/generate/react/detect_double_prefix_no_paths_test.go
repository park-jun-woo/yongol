//ff:func feature=gen-react type=test control=sequence
//ff:what detectDoublePrefix 빈 path 집합은 이중 접두가 아니므로 nil (BUG-110)

package react

import "testing"

func TestDetectDoublePrefix_NoPaths_Nil(t *testing.T) {
	if err := detectDoublePrefix("/api", nil); err != nil {
		t.Fatalf("expected nil for empty path set, got %v", err)
	}
}
