//ff:func feature=migration type=accessor control=sequence
//ff:what DropColumn.Description — 파일명/헤더용 설명 문자열
package migration

func (op DropColumn) Description() string { return "drop column " + op.Table + "." + op.Column }
