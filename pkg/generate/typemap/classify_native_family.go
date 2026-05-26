//ff:func feature=gen-typemap type=util control=sequence
//ff:what classifyNativeFamily — native family head 토큰을 PGFamily 로 매핑

package typemap

// classifyNativeFamily maps native-family head tokens to their PGFamily.
func classifyNativeFamily(head string) (PGFamily, bool) {
	if integerHeads[head] {
		return FamilyInteger, true
	}
	if floatHeads[head] {
		return FamilyFloat, true
	}
	if stringHeads[head] {
		return FamilyString, true
	}
	if booleanHeads[head] {
		return FamilyBoolean, true
	}
	return FamilyUnsupported, false
}
