//ff:func feature=migration type=accessor control=sequence
//ff:what AddColumn.Description — 파일명/헤더용 설명 문자열
package migration

func (op AddColumn) Description() string { return "add column " + op.Table + "." + op.Column.Name }
