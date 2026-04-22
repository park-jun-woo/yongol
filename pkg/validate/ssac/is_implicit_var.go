//ff:func feature=validate type=util control=selection topic=ssac-structural
//ff:what isImplicitVar — 암시 소스(request/currentUser/query/message) 판정

package ssac

// isImplicitVar reports whether name is one of the implicitly-available
// sources. HTTP funcs implicitly have {request, currentUser, query};
// subscribe funcs additionally have {message}. We accept all four for
// both flavors here and let the per-rule logic forbid them where
// appropriate.
func isImplicitVar(name string) bool {
	switch name {
	case "request", "currentUser", "query", "message":
		return true
	}
	return false
}
