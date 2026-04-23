//ff:func feature=gen-gogin type=util control=sequence topic=csrf
//ff:what csrfAlwaysInactive — csrf 비활성 MainBlock 용 고정 predicate

package boot

import "github.com/park-jun-woo/yongol/pkg/yongol"

// csrfAlwaysInactive is the Active predicate for the dormant csrf
// placeholder block emitted when hasCsrf=false. Returning false lets
// collectActiveBlocks filter the block out without any state read.
func csrfAlwaysInactive(fs *yongol.Fullstack) bool {
	_ = fs
	return false
}
