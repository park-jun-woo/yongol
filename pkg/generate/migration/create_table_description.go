//ff:func feature=migration type=accessor control=sequence
//ff:what CreateTable.Description — 파일명/헤더용 설명 문자열
package migration

func (op CreateTable) Description() string { return "create table " + op.Table.Name }
