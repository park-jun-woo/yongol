//ff:func feature=gen-fastapi type=adapter control=sequence
//ff:what CheckEnum — elementHeadMeta 는 항상 nil 반환 (element head 에 enum 없음)

package types

// CheckEnum always returns nil because a bare element head has no CHECK
// constraint.
func (m elementHeadMeta) CheckEnum() []string { return nil }
