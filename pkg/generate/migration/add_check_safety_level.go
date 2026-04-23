//ff:func feature=migration type=accessor control=sequence
//ff:what AddCheck.SafetyLevel — 항상 SafetySafe
package migration

func (op AddCheck) SafetyLevel() SafetyLevel { return SafetySafe }
