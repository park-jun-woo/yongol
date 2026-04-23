//ff:func feature=migration type=accessor control=sequence
//ff:what DropForeignKey.Description — 파일명/헤더용 설명 문자열
package migration

func (op DropForeignKey) Description() string { return "drop FK " + op.Name }
