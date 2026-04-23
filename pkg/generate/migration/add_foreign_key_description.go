//ff:func feature=migration type=accessor control=sequence
//ff:what AddForeignKey.Description — 파일명/헤더용 설명 문자열
package migration

func (op AddForeignKey) Description() string { return "add FK " + op.FK.Name }
