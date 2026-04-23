//ff:func feature=migration type=accessor control=sequence
//ff:what RenameTable.Description — 파일명/헤더용 설명 문자열
package migration

func (op RenameTable) Description() string { return "rename table " + op.From + " → " + op.To }
