//ff:func feature=migration type=accessor control=sequence
//ff:what DropForeignKey.SafetyLevel — 항상 SafetySafe
package migration

func (op DropForeignKey) SafetyLevel() SafetyLevel { return SafetySafe }
