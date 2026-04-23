//ff:func feature=migration type=accessor control=selection
//ff:what DropColumn.SafetyLevel — @allow_destructive 없으면 SafetyWarning
package migration

func (op DropColumn) SafetyLevel() SafetyLevel {
	if op.AllowDestructive {
		return SafetySafe
	}
	return SafetyWarning
}
