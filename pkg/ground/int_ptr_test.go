//ff:func feature=rule type=test-helper control=sequence
//ff:what intPtr — int → *int 테스트 헬퍼 (*FieldConstraint.MaxLength 세팅용)

package ground

// intPtr returns a pointer to an int (handy for *FieldConstraint.MaxLength).
func intPtr(n int) *int { return &n }
