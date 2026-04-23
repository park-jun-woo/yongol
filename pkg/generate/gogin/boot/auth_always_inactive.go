//ff:func feature=gen-gogin type=util control=sequence
//ff:what authAlwaysInactive — auth 비활성 placeholder MainBlock 용 predicate

package boot

import "github.com/park-jun-woo/yongol/pkg/yongol"

// authAlwaysInactive is the Active predicate for the dormant auth-init
// placeholder block emitted when prepared.Auth.Present=false.
func authAlwaysInactive(fs *yongol.Fullstack) bool {
	_ = fs
	return false
}
