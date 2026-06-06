//ff:func feature=gen-react type=test control=sequence
//ff:what detectDoublePrefix 모든 path 가 baseUrl 접두를 공유하면 이중 접두 에러 (BUG-110)

package react

import "testing"

func TestDetectDoublePrefix_AllSharePrefix_Error(t *testing.T) {
	paths := []string{"/api/health", "/api/v1/admin/buildings"}
	if err := detectDoublePrefix("/api", paths); err == nil {
		t.Fatal("expected error when every path shares the baseUrl prefix")
	}
}
