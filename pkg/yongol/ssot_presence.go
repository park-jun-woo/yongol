//ff:type feature=orchestrator type=model
//ff:what SSOTPresence — SSOT 디렉토리/파일 감지 3-상태 enum

package yongol

// SSOTPresence captures the declaration state of an SSOT. Supersedes the
// prior binary "detected / not detected" distinction so validators can tell
// apart "user opted out" from "user declared the SSOT but left it empty".
type SSOTPresence int

const (
	// SSOTAbsent — neither directory nor file present. User opts out.
	SSOTAbsent SSOTPresence = iota
	// SSOTDeclared — directory exists but no matching content file found.
	// User signaled intent but hasn't populated (WIP / mistake).
	SSOTDeclared
	// SSOTPopulated — directory exists and at least one content file present.
	SSOTPopulated
)
