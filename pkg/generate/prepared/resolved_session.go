//ff:type feature=generate type=model
//ff:what Session — session 백엔드 파생 상태 (non-zero 보장)

package prepared

// Session carries the derived session backend configuration used by
// codegen. Present only when manifest declares session.backend OR SSaC
// uses session.* calls; otherwise State.ActiveBackends.Session is nil
// and no session blocks are emitted.
//
// Backend is one of "postgres" or "memory"; never empty. prepared.New
// applies the "memory" default so emitters can branch on a known value
// set.
type Session struct {
	Backend string
}
