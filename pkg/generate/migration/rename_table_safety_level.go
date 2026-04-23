//ff:func feature=migration type=accessor control=sequence
//ff:what RenameTable.SafetyLevel — 항상 SafetySafe
package migration

func (op RenameTable) SafetyLevel() SafetyLevel { return SafetySafe }
