//ff:func feature=gen-hurl type=test-helper control=sequence
//ff:what pathLooksLikeSignup — auth signup 경로 판정

package hurl

// pathLooksLikeSignup reports whether p is one of the canonical signup
// paths used by the auth-shape inference helpers.
func pathLooksLikeSignup(p string) bool {
	return p == "/auth/register" || p == "/auth/signup" || p == "/auth/join"
}
