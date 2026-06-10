//ff:func feature=gen-react type=generator control=sequence
//ff:what renderLogoutHandler — 모드별 로그아웃 핸들러 방출 (op 호출 → store clear(bearer) → /login)

package react

import (
	"fmt"
	"strings"
)

// renderLogoutHandler renders the handleLogout function body of a layout
// component (page-flow Phase010). Mode wiring mirrors the auth-flow
// derivation (prepared.AuthFor — the caller passes the resolved mode):
//   - bearer: the optional server op is called best-effort (a failed
//     server logout must not trap the user in the session), then the
//     session store is cleared and the user lands on /login.
//   - cookie: the server op call *is* the logout — the session lives in
//     an httpOnly cookie only the server can end (a valueless
//     data-logout in cookie mode draws TM-38) — then navigate to /login.
//
// The /login destination reuses the existing convention
// (clearSessionAndLogin / writeCookie401Redirect); promoting it to an
// SSOT decision is a recorded follow-up outside the page-flow series.
func renderLogoutHandler(operationID, authMode string) string {
	var sb strings.Builder
	if operationID != "" {
		sb.WriteString("  const handleLogout = async () => {\n")
		fmt.Fprintf(&sb, "    await api.%s().catch(() => {}) // server logout is best-effort\n", operationID)
	} else {
		sb.WriteString("  const handleLogout = () => {\n")
	}
	if authMode == "bearer" {
		sb.WriteString("    useAuthStore.getState().clear()\n")
	}
	sb.WriteString("    navigate('/login')\n")
	sb.WriteString("  }\n")
	return sb.String()
}
