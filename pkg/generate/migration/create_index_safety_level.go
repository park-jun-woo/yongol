//ff:func feature=migration type=accessor control=sequence
//ff:what CreateIndex.SafetyLevel — 항상 SafetySafe
package migration

func (op CreateIndex) SafetyLevel() SafetyLevel { return SafetySafe }
