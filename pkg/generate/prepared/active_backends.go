//ff:type feature=generate type=model
//ff:what ActiveBackends — 활성 백엔드 (nil = 비활성) 모음

package prepared

// ActiveBackends collects the four optional runtime backends whose
// activation is driven by manifest declaration OR SSaC usage. Each field
// is nil when the corresponding subsystem is not in use; emitters branch
// on nil instead of parsing raw manifest paths.
//
// Invariant: Session != nil iff (manifest.session.backend set) OR
// (SSaC issues @call session.*). Same shape for Cache, File, Queue.
type ActiveBackends struct {
	Session *Session
	Cache   *Cache
	File    *File
	Queue   *Queue
}
