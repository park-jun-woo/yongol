//ff:func feature=validate type=util control=sequence
//ff:what WithGround — 사전 조립된 *rule.Ground를 주입 (테스트/세밀 제어용)
package validate

import "github.com/park-jun-woo/yongol/pkg/rule"

// WithGround injects a pre-built Ground. When omitted, Validate builds one
// via pkg/ground.Build(fs).
func WithGround(g *rule.Ground) Option {
	return func(c *config) { c.ground = g }
}
