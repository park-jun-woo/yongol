//ff:func feature=migration type=accessor control=sequence
//ff:what AlterColumnNullable.Destructive — NOT NULL 추가(To=false)만 파괴적
package migration

func (op AlterColumnNullable) Destructive() bool { return !op.To }
