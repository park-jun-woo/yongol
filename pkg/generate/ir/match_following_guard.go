//ff:func feature=gen-ir type=util control=sequence
//ff:what matchFollowingGuard -- @empty/@exists 가드가 대상 변수에 매칭되는지 판정

package ir

// matchFollowingGuard checks whether next is an @empty or @exists guard
// targeting varName. Returns the matching OpKind or zero (OpGet) if no match.
func matchFollowingGuard(next Op, varName string) OpKind {
	if next.Kind == OpEmpty && next.Empty != nil && next.Empty.VarName == varName {
		return OpEmpty
	}
	if next.Kind == OpExists && next.Exists != nil && next.Exists.VarName == varName {
		return OpExists
	}
	return OpGet
}
