//ff:func feature=validate type=accessor control=iteration dimension=1
//ff:what step.isPresent — Kinds 전부가 Fullstack에 존재하면 true
package validate

import "github.com/park-jun-woo/yongol/pkg/yongol"

func (s step) isPresent(fs *yongol.Fullstack) bool {
	for _, k := range s.Kinds {
		if !kindPresent(fs, k) {
			return false
		}
	}
	return true
}
