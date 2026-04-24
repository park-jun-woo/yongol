//ff:func feature=gen-hurl type=test-helper control=sequence
//ff:what pathLooksLikeLogin — auth login 경로 판정

package hurl

// pathLooksLikeLogin reports whether p is one of the canonical login
// paths used by the auth-shape inference helpers.
func pathLooksLikeLogin(p string) bool {
	return p == "/auth/login" || p == "/auth/signin"
}
