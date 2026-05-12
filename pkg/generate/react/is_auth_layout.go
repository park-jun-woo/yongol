//ff:func feature=gen-react type=util control=sequence
//ff:what 레이아웃 이름이 인증 레이아웃인지 확인한다

package react

// isAuthLayout returns true if the layout name is "auth".
// Auth layouts host public pages (login, register) and are never wrapped
// with ProtectedRoute.
func isAuthLayout(name string) bool {
	return name == "auth"
}
