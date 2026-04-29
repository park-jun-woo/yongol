//ff:func feature=gen-gogin type=util control=selection
//ff:what mapNativeFamily — head 토큰이 native family (Integer/Float/String/Boolean) 면 binding 반환

package types

// mapNativeFamily routes Integer / Float / String / Boolean head tokens
// to their native bindings (or pointer variants for nullable). Returns
// ok=false when head matches no native family — caller treats it as
// unsupported.
func mapNativeFamily(head string, notNull bool, defaultLiteral string) (GoTypeBinding, bool) {
	switch {
	case isIntegerHead(head):
		return nativeInteger(notNull, defaultLiteral), true
	case isFloatHead(head):
		return nativeFloat(notNull, defaultLiteral), true
	case isStringHead(head):
		return nativeString(notNull, defaultLiteral), true
	case isBooleanHead(head):
		return nativeBoolean(notNull, defaultLiteral), true
	}
	return GoTypeBinding{}, false
}
