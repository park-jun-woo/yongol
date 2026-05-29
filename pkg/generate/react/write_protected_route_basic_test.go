//ff:func feature=gen-react type=test control=sequence
//ff:what writeProtectedRoute ProtectedRoute.tsx 생성 검증

package react

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteProtectedRoute(t *testing.T) {
	dir := t.TempDir()
	if err := writeProtectedRoute(dir); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "components", "ProtectedRoute.tsx"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)

	assertContains(t, content, "import { Navigate } from 'react-router-dom'")
	assertContains(t, content, "export default function ProtectedRoute")
	assertContains(t, content, "localStorage.getItem('access_token')")
	assertContains(t, content, `<Navigate to="/login" replace />`)
}
