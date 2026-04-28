//ff:func feature=gen-gogin type=util control=selection
//ff:what resolveErrStatus — 시퀀스 타입별 기본 에러 HTTP status 결정

package ssac

// resolveErrStatus returns the effective HTTP status for a guard sequence.
// Priority: explicit ErrStatus > type default.
func resolveErrStatus(seqType string, explicitStatus int) int {
	if explicitStatus != 0 {
		return explicitStatus
	}
	switch seqType {
	case "empty":
		return 404
	case "exists", "state":
		return 409
	case "auth":
		return 403
	case "call":
		return 500
	case "eval":
		// SSaC enforces an explicit STATUS via S-68. The 500 fallback only
		// fires when callers reach codegen with ErrStatus 0 (e.g. snapshot
		// tests), keeping emission deterministic without papering over the
		// SSaC error.
		return 500
	default:
		return 500
	}
}
