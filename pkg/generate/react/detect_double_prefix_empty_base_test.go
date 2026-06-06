//ff:func feature=gen-react type=test control=sequence
//ff:what detectDoublePrefix 빈 baseUrl 은 이중 접두가 아니므로 nil (BUG-110)

package react

import "testing"

func TestDetectDoublePrefix_EmptyBaseURL_Nil(t *testing.T) {
	paths := []string{"/api/health", "/api/v1/admin/buildings"}
	if err := detectDoublePrefix("", paths); err != nil {
		t.Fatalf("expected nil for empty baseUrl, got %v", err)
	}
}
