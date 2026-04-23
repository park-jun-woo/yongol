//ff:func feature=migration type=accessor control=sequence
//ff:what DropIndex.Description — 파일명/헤더용 설명 문자열
package migration

func (op DropIndex) Description() string { return "drop index " + op.Name }
