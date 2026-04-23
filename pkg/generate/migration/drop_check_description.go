//ff:func feature=migration type=accessor control=sequence
//ff:what DropCheck.Description — 파일명/헤더용 설명 문자열
package migration

func (op DropCheck) Description() string { return "drop check " + op.Name }
