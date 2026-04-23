//ff:func feature=migration type=accessor control=sequence
//ff:what AlterColumnType.Description — 헤더용 from→to 타입 표기
package migration

import "fmt"

func (op AlterColumnType) Description() string {
	return fmt.Sprintf("alter %s.%s type %s→%s", op.Table, op.Column, op.From.SQL(), op.To.SQL())
}
