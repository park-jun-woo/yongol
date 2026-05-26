//ff:func feature=gen-ir type=util control=sequence
//ff:what rateLimitHasEntries -- manifest.backend.rate_limit 항목 존재 여부

package ir

import "github.com/park-jun-woo/yongol/pkg/yongol"

// rateLimitHasEntries returns true when manifest.backend.rate_limit has
// at least one entry.
func rateLimitHasEntries(fs *yongol.Fullstack) bool {
	if fs == nil || fs.Manifest == nil {
		return false
	}
	return len(fs.Manifest.Backend.RateLimit) > 0
}
