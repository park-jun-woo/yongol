//ff:func feature=generate type=util control=sequence
//ff:what queueBackendFor — manifest + SSaC 사용 여부로 queue 활성 판정 및 기본값 해석

package prepared

import "github.com/park-jun-woo/yongol/pkg/yongol"

// queueBackendFor returns non-nil iff the queue subsystem is in use.
// Activation triggers: manifest.queue.backend set, or any SSaC function
// with @subscribe / publish sequence. Default "postgres" matches the
// behavior that block_queue_init previously applied when manifest was
// silent.
func queueBackendFor(fs *yongol.Fullstack) *Queue {
	if manifestDeclaresQueue(fs) {
		return &Queue{Backend: fs.Manifest.Queue.Backend}
	}
	if ssacUsesQueue(fs) {
		return &Queue{Backend: "postgres"}
	}
	return nil
}
