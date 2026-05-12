//ff:func feature=gen-react type=generator control=sequence
//ff:what ProtectedRoute.tsx — JWT 미존재 시 /login 리다이렉트 가드 컴포넌트 방출

package react

import (
	"os"
	"path/filepath"
)

// writeProtectedRoute emits src/components/ProtectedRoute.tsx.
// The component checks localStorage for an access_token and redirects
// unauthenticated users to /login.
func writeProtectedRoute(srcDir string) error {
	dir := filepath.Join(srcDir, "components")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}

	const src = `import { Navigate } from 'react-router-dom'

export default function ProtectedRoute({ children }: { children: React.ReactNode }) {
  const token = localStorage.getItem('access_token')
  if (!token) return <Navigate to="/login" replace />
  return <>{children}</>
}
`
	return os.WriteFile(filepath.Join(dir, "ProtectedRoute.tsx"), []byte(src), 0o644)
}
