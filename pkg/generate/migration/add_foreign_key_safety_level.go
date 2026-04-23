//ff:func feature=migration type=accessor control=sequence
//ff:what AddForeignKey.SafetyLevel — 항상 SafetySafe
package migration

func (op AddForeignKey) SafetyLevel() SafetyLevel { return SafetySafe }
