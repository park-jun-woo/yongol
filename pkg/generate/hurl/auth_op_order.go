//ff:func feature=gen-hurl type=util control=sequence
//ff:what authOpOrder — smoke 에 배치할 인증 operation 고정 순서 (Register → Login)

package hurl

// authOpOrder returns the fixed lookup order used by buildAuthPair.
// The order is deliberately Register-first (creates the account) then
// Login (verifies the just-created credentials) — the only ordering
// that works on an empty DB. See BUG-015.
func authOpOrder() []string {
	return []string{"Register", "Login"}
}
