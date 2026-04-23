//ff:func feature=migration type=accessor control=sequence
//ff:what CreateTable.SafetyLevel — 항상 SafetySafe
package migration

func (op CreateTable) SafetyLevel() SafetyLevel { return SafetySafe }
