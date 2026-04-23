//ff:func feature=migration type=accessor control=sequence
//ff:what AlterColumnDefault.SafetyLevel — 항상 SafetySafe
package migration

func (op AlterColumnDefault) SafetyLevel() SafetyLevel { return SafetySafe }
