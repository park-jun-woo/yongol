//ff:type feature=policy type=model
//ff:what OwnershipMapping — @ownership 어노테이션에서 추출한 소유권 매핑
package rego

// OwnershipMapping represents a @ownership annotation.
type OwnershipMapping struct {
	Resource   string
	Table      string
	Column     string
	JoinTable  string
	JoinFK     string
	SourceLine int // 1-based line number of the @ownership comment in the source .rego (0 = unknown)
}
