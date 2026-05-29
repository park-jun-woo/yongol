//ff:func feature=rule type=test-helper control=sequence
//ff:what withResponseConstraints — pre-built response constraint map 을 Fullstack 에 부착

package ground

import (
	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// withResponseConstraints attaches pre-built response constraint maps.
func withResponseConstraints(m map[string]map[string]oapiparser.FieldConstraint) func(*yongol.Fullstack) {
	return func(fs *yongol.Fullstack) { fs.ResponseConstraints = m }
}
