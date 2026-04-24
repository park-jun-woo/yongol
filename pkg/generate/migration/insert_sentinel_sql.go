//ff:func feature=migration type=accessor control=sequence
//ff:what InsertSentinel.SQL — preserved INSERT body 를 그대로 반환
package migration

// SQL returns the raw INSERT statement text as authored in the DDL
// file (including the terminating `;`).
func (op InsertSentinel) SQL() string { return op.Body }
