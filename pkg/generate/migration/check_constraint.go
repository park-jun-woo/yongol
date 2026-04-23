//ff:type feature=migration type=model
//ff:what CheckConstraint — CHECK (...) 제약 정의
package migration

// CheckConstraint describes a CHECK (...) constraint.
type CheckConstraint struct {
	Name       string // auto-generated when user omitted
	Expression string // canonical expression text
}
