//ff:func feature=migration type=accessor control=selection
//ff:what AlterColumnType.SafetyLevel — USING 없으면 SafetyWarning
package migration

func (op AlterColumnType) SafetyLevel() SafetyLevel {
	if op.Using != "" {
		return SafetySafe
	}
	return SafetyWarning
}
