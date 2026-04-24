//ff:func feature=migration type=accessor control=sequence
//ff:what InsertSentinel.Description — 파일명/헤더용 설명 문자열
package migration

// Description returns a short human-readable label used in migration
// file headers.
func (op InsertSentinel) Description() string { return "insert sentinel " + op.Table }
