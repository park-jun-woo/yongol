//ff:func feature=migration type=accessor control=selection
//ff:what AlterColumnNullable.SafetyLevel — NOT NULL 추가 + backfill 없으면 SafetyError
package migration

func (op AlterColumnNullable) SafetyLevel() SafetyLevel {
	if !op.To && op.Backfill == "" {
		return SafetyError
	}
	return SafetySafe
}
