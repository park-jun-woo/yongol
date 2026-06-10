//ff:func feature=gen-react type=test control=sequence
//ff:what writeProtectedRoute cookie 모드 — 낙관 통과 가드(store/redirect 없음) 생성 검증

package react

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteProtectedRoute_CookieOptimistic(t *testing.T) {
	dir := t.TempDir()
	if err := writeProtectedRoute(dir, false); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "components", "ProtectedRoute.tsx"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)

	assertContains(t, content, "export default function ProtectedRoute")
	assertContains(t, content, "return <>{children}</>")
	// the httpOnly session is unreadable — no token check, no redirect
	assertNotContains(t, content, "useAuthStore")
	assertNotContains(t, content, "Navigate")
	assertNotContains(t, content, "localStorage")
}
