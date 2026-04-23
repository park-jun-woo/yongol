//ff:func feature=generate type=util control=sequence
//ff:what sessionBackendFor — manifest + SSaC 사용 여부로 session 활성 판정 및 기본값 해석

package prepared

import "github.com/park-jun-woo/yongol/pkg/yongol"

// sessionBackendFor returns non-nil iff the session subsystem is in
// use — either manifest declares session.backend or some SSaC service
// func issues an @call session.*. The Backend field carries the
// resolved default ("memory" when manifest is absent but SSaC uses
// session).
//
// This is the single source of truth for "is session active in the
// generated backend?". Emitters must not consult fs.Manifest.Session
// directly.
func sessionBackendFor(fs *yongol.Fullstack) *Session {
	if manifestDeclaresSession(fs) {
		return &Session{Backend: fs.Manifest.Session.Backend}
	}
	if ssacUsesSessionCalls(fs) {
		return &Session{Backend: "memory"}
	}
	return nil
}
