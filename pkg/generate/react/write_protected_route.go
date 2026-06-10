//ff:func feature=gen-react type=generator control=sequence
//ff:what ProtectedRoute.tsx — bearer는 세션 store 토큰 판정, cookie는 낙관 통과 가드 방출

package react

import (
	"os"
	"path/filepath"
)

// writeProtectedRoute emits src/components/ProtectedRoute.tsx, the per-page
// guard App.tsx wraps protected routes with (Phase005).
//
//   - bearer: reads the token from the Phase003 session store
//     (src/stores/auth.ts — the same store the API client injects from; no
//     more direct localStorage key reads) and redirects to /login when
//     absent.
//   - cookie: the session lives in an httpOnly cookie the client cannot
//     read, so the guard passes optimistically; a 401 from any protected
//     fetch is converged to /login by the API client (Phase004 semantics —
//     no extra session-check endpoint is forced on the SSOT).
func writeProtectedRoute(srcDir string, bearer bool) error {
	dir := filepath.Join(srcDir, "components")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}

	const bearerSrc = `import { Navigate } from 'react-router-dom'
import { useAuthStore } from '../stores/auth'

export default function ProtectedRoute({ children }: { children: React.ReactNode }) {
  const token = useAuthStore((s) => s.token)
  if (!token) return <Navigate to="/login" replace />
  return <>{children}</>
}
`
	const cookieSrc = `// Cookie-mode guard: the session is an httpOnly cookie the client cannot
// read, so the guard passes optimistically — a 401 from any protected fetch
// is converged to /login by the API client.
export default function ProtectedRoute({ children }: { children: React.ReactNode }) {
  return <>{children}</>
}
`
	src := cookieSrc
	if bearer {
		src = bearerSrc
	}
	return os.WriteFile(filepath.Join(dir, "ProtectedRoute.tsx"), []byte(src), 0o644)
}
