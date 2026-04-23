//ff:func feature=migration type=accessor control=sequence
//ff:what AlterColumnDefault.Description — 헤더용 표기
package migration

func (op AlterColumnDefault) Description() string {
	return "alter default " + op.Table + "." + op.Column + " → " + op.To
}
