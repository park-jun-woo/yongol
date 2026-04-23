//ff:func feature=migration type=accessor control=sequence
//ff:what DropIndex.SafetyLevel — 항상 SafetySafe
package migration

func (op DropIndex) SafetyLevel() SafetyLevel { return SafetySafe }
