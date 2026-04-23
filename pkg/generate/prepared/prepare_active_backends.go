//ff:func feature=generate type=util control=sequence
//ff:what activeBackendsFor — session/cache/file/queue 활성 상태를 한 번에 계산

package prepared

import "github.com/park-jun-woo/yongol/pkg/yongol"

// activeBackendsFor computes the nil-or-resolved shape for every
// optional runtime backend. Each field follows the same invariant:
// non-nil iff manifest declares the backend OR SSaC uses it.
func activeBackendsFor(fs *yongol.Fullstack) ActiveBackends {
	return ActiveBackends{
		Session: sessionBackendFor(fs),
		Cache:   cacheBackendFor(fs),
		File:    fileBackendFor(fs),
		Queue:   queueBackendFor(fs),
	}
}
