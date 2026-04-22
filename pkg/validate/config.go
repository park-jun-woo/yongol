//ff:type feature=validate type=model
//ff:what config — Validate 내부 설정 컨테이너 (ground)
package validate

import (
	"github.com/park-jun-woo/yongol/pkg/rule"
)

// config holds runtime configuration applied by Options.
type config struct {
	ground  *rule.Ground
	artsDir string
}
