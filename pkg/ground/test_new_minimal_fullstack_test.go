//ff:func feature=rule type=test-helper control=iteration dimension=1
//ff:what newMinimalFullstack — opts 를 순회 적용해 *yongol.Fullstack 을 조립

package ground

import (
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// newMinimalFullstack returns an empty *yongol.Fullstack with nil SSOT fields.
// Opts mutate individual fields.
func newMinimalFullstack(opts ...func(*yongol.Fullstack)) *yongol.Fullstack {
	fs := &yongol.Fullstack{}
	for _, o := range opts {
		o(fs)
	}
	return fs
}
