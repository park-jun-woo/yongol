//ff:func feature=migration type=accessor control=selection
//ff:what AddColumn.SafetyLevel — NOT NULL + default/backfill 없으면 SafetyError
package migration

func (op AddColumn) SafetyLevel() SafetyLevel {
	if !op.Column.Nullable && op.Column.Default == "" && op.Backfill == "" {
		return SafetyError
	}
	return SafetySafe
}
