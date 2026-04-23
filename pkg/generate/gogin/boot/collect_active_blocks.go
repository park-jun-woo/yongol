//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what collectActiveBlocks — Fullstack 상태에 따라 활성 블록만 수집

package boot

import (
	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// collectActiveBlocks assembles blocks in fixed order and filters to those
// whose Active condition passes (nil Active = always active).
//
// Blocks whose activation has been migrated to prepared.State (session,
// cache, file, queue in Phase001) are appended only when the matching
// ActiveBackends field is non-nil — their constructors take a resolved
// value and cannot panic on missing manifest subtrees.
func collectActiveBlocks(fs *yongol.Fullstack, p prepared.State, modulePath string) []MainBlock {
	candidates := baseCandidateBlocks(fs, p, modulePath)
	var active []MainBlock
	for _, b := range candidates {
		if b.Active == nil || b.Active(fs) {
			active = append(active, b)
		}
	}
	return active
}
