//ff:func feature=migration type=accessor control=sequence
//ff:what AddCheck.Description — 파일명/헤더용 설명 문자열
package migration

func (op AddCheck) Description() string { return "add check " + op.Check.Name }
