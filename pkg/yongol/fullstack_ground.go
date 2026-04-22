//ff:func feature=orchestrator type=accessor control=sequence
//ff:what Fullstack.Ground — 적재된 *rule.Ground 반환 (nil이면 SetGround 선행 필요)
package yongol

import "github.com/park-jun-woo/yongol/pkg/rule"

// Ground returns the rule.Ground associated with this Fullstack.
// Returns nil until SetGround is called. pkg/validate.Validate calls
// SetGround(ground.Build(fs)) at entry so downstream step functions can
// consume it through Fullstack without a separate parameter.
func (fs *Fullstack) Ground() *rule.Ground {
	return fs.ground
}
