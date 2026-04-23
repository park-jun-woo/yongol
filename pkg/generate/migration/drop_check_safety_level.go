//ff:func feature=migration type=accessor control=sequence
//ff:what DropCheck.SafetyLevel — 항상 SafetySafe
package migration

func (op DropCheck) SafetyLevel() SafetyLevel { return SafetySafe }
