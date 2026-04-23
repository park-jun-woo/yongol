//ff:func feature=migration type=accessor control=sequence
//ff:what CreateIndex.Description — 파일명/헤더용 설명 문자열
package migration

func (op CreateIndex) Description() string { return "create index " + op.Index.Name }
