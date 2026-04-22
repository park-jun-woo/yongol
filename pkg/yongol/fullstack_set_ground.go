//ff:func feature=orchestrator type=accessor control=sequence
//ff:what Fullstack.SetGround — *rule.Ground를 Fullstack에 바인딩 (cycle 회피용 외부 주입)
package yongol

import "github.com/park-jun-woo/yongol/pkg/rule"

// SetGround binds a rule.Ground to this Fullstack. pkg/yongol cannot import
// pkg/ground (circular), so the caller (pkg/validate.Validate) constructs
// Ground via pkg/ground.Build(fs) and injects it here.
func (fs *Fullstack) SetGround(g *rule.Ground) {
	fs.ground = g
}
