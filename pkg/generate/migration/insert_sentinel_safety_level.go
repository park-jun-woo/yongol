//ff:func feature=migration type=accessor control=sequence
//ff:what InsertSentinel.SafetyLevel — 항상 SafetySafe
package migration

// SafetyLevel is SafetySafe — sentinel INSERTs introduce data but are
// idempotent and required for FK DEFAULT 0 patterns.
func (op InsertSentinel) SafetyLevel() SafetyLevel { return SafetySafe }
