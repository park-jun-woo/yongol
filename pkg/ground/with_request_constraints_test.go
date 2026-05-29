//ff:func feature=rule type=test-helper control=sequence
//ff:what withRequestConstraints — pre-built request constraint map 을 Fullstack 에 부착

package ground

import (
	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// withRequestConstraints attaches pre-built request constraint maps.
func withRequestConstraints(m map[string]map[string]oapiparser.FieldConstraint) func(*yongol.Fullstack) {
	return func(fs *yongol.Fullstack) { fs.RequestConstraints = m }
}
