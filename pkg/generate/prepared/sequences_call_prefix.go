//ff:func feature=generate type=util control=iteration dimension=1
//ff:what sequencesCallPrefix — SSaC 시퀀스 목록에 prefix로 시작하는 @call이 있는지

package prepared

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// sequencesCallPrefix returns true when any call-type sequence has a
// Model name starting with the given prefix (for example "session." or
// "cache."). Shared across prepareSessionBackend / prepareCacheBackend /
// prepareFileBackend / prepareQueueBackend to keep usage detection in
// one place.
func sequencesCallPrefix(seqs []ssac.Sequence, prefix string) bool {
	for _, seq := range seqs {
		if seq.Type == "call" && strings.HasPrefix(seq.Model, prefix) {
			return true
		}
	}
	return false
}
