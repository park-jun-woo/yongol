//ff:func feature=migration type=accessor control=sequence
//ff:what RenameColumn.SafetyLevel — 항상 SafetySafe
package migration

func (op RenameColumn) SafetyLevel() SafetyLevel { return SafetySafe }
