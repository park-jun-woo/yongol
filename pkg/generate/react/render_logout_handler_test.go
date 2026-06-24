//ff:func feature=gen-react type=test control=sequence
//ff:what renderLogoutHandler — bearer/cookie × op 유무 4분기 핸들러 라인 검증

package react

import "testing"

func TestRenderLogoutHandler(t *testing.T) {
	t.Run("bearer with op", func(t *testing.T) {
		out := renderLogoutHandler("Logout", "bearer", "")
		assertContains(t, out, "const handleLogout = async () => {")
		assertContains(t, out, "await api.Logout().catch(() => {})")
		assertContains(t, out, "useAuthStore.getState().clear()")
		assertContains(t, out, "navigate('/login')")
	})

	t.Run("bearer valueless", func(t *testing.T) {
		out := renderLogoutHandler("", "bearer", "")
		assertContains(t, out, "const handleLogout = () => {")
		assertNotContains(t, out, "api.")
		assertContains(t, out, "useAuthStore.getState().clear()")
		assertContains(t, out, "navigate('/login')")
	})

	t.Run("cookie with op", func(t *testing.T) {
		out := renderLogoutHandler("Logout", "cookie", "")
		assertContains(t, out, "await api.Logout().catch(() => {})")
		assertNotContains(t, out, "useAuthStore")
		assertContains(t, out, "navigate('/login')")
	})

	t.Run("cookie valueless", func(t *testing.T) {
		out := renderLogoutHandler("", "cookie", "")
		assertContains(t, out, "const handleLogout = () => {")
		assertNotContains(t, out, "api.")
		assertNotContains(t, out, "useAuthStore")
		assertContains(t, out, "navigate('/login')")
	})

	t.Run("bearer with op and refresh body key", func(t *testing.T) {
		out := renderLogoutHandler("Logout", "bearer", "refresh_token")
		assertContains(t, out, "const handleLogout = async () => {")
		assertContains(t, out, `await api.Logout({ refresh_token: useAuthStore.getState().refresh ?? '' }).catch(() => {})`)
		assertContains(t, out, "useAuthStore.getState().clear()")
		assertContains(t, out, "navigate('/login')")
	})
}
