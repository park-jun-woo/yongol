//ff:func feature=orchestrator type=accessor control=sequence
//ff:what PresenceOf — SSOTKind → Fullstack.Presences 조회 (미등록 시 Absent)

package yongol

// PresenceOf returns the SSOTPresence for the given kind. Kinds with no entry
// in fs.Presences are treated as SSOTAbsent (never detected).
func (fs *Fullstack) PresenceOf(k SSOTKind) SSOTPresence {
	if fs == nil || fs.Presences == nil {
		return SSOTAbsent
	}
	return fs.Presences[k]
}
