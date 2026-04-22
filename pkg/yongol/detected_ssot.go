//ff:type feature=orchestrator type=model
//ff:what 탐지된 SSOT 파일/디렉토리 정보
package yongol

// DetectedSSOT represents a detected SSOT file or directory. Presence captures
// whether the directory merely exists (SSOTDeclared) or contains matching
// content (SSOTPopulated). SSOTAbsent SSOTs are not emitted as DetectedSSOT.
type DetectedSSOT struct {
	Kind     SSOTKind
	Path     string // absolute path to the relevant directory or file
	Presence SSOTPresence
}
