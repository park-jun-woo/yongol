//ff:func feature=cli type=util control=iteration dimension=1
//ff:what countPreserved — splits preserved path slice into files with and without a reason comment

package main

import (
	"github.com/park-jun-woo/yongol/pkg/contract"
)

// countPreserved returns (withReason, withoutReason) counts. A preserve
// reason is any non-empty string returned by contract.ParsePreserveReason.
// Read errors are treated as "no reason" so a best-effort summary is
// produced even when a file becomes unreadable between CollectPreserved
// and this pass.
func countPreserved(paths []string) (withReason, withoutReason int) {
	for _, p := range paths {
		reason, err := contract.ParsePreserveReason(p)
		if err != nil || reason == "" {
			withoutReason++
			continue
		}
		withReason++
	}
	return withReason, withoutReason
}
