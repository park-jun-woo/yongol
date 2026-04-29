//ff:func feature=gen-gogin type=util control=sequence
//ff:what unsupportedBinding — 다중 토큰 / CREATE TYPE / 미지원 PG 타입의 거절 binding

package types

// unsupportedBinding returns a sentinel binding that fails fast at the
// emit site. validate D-11 fires for the same column before generate
// runs, so this struct is mostly defensive — but if a user bypasses
// validate (e.g. invokes generate directly) the empty templates surface
// as a Go compile error rather than a silent miscompile.
//
// reason is a short human-readable explanation surfaced via the
// SqlcGoType field for diagnostics; emit sites should not interpret it
// programmatically.
func unsupportedBinding(reason string) GoTypeBinding {
	return GoTypeBinding{
		SqlcGoType: "/* unsupported: " + reason + " */",
		ApiField:   "/* unsupported: " + reason + " */",
		Kind:       KindUnsupported,
		Supported:  false,
	}
}
