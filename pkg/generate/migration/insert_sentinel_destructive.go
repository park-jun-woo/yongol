//ff:func feature=migration type=accessor control=sequence
//ff:what InsertSentinel.Destructive — 항상 false (ON CONFLICT DO NOTHING 가 재실행 안전)
package migration

// Destructive is always false — sentinel INSERTs are idempotent by
// construction (ON CONFLICT DO NOTHING enforced by D-10).
func (op InsertSentinel) Destructive() bool { return false }
