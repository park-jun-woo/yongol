//ff:func feature=migration type=accessor control=sequence
//ff:what DropTable.Description — 파일명/헤더용 설명 문자열
package migration

func (op DropTable) Description() string { return "drop table " + op.Name }
