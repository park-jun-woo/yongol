//ff:func feature=migration type=accessor control=selection
//ff:what AlterColumnNullable.Description — To=true 일 때 drop / 아니면 set
package migration

func (op AlterColumnNullable) Description() string {
	if op.To {
		return "drop NOT NULL on " + op.Table + "." + op.Column
	}
	return "set NOT NULL on " + op.Table + "." + op.Column
}
