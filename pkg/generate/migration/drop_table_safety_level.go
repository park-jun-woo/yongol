//ff:func feature=migration type=accessor control=selection
//ff:what DropTable.SafetyLevel — @allow_destructive 없으면 SafetyWarning
package migration

func (op DropTable) SafetyLevel() SafetyLevel {
	if op.AllowDestructive {
		return SafetySafe
	}
	return SafetyWarning
}
