//ff:func feature=validate type=test control=sequence topic=states
//ff:what XDM-27 test — stateDiagram <entity>_<field> → DDL column 존재 검증
package ddl_statemachine

import (
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func fsWithGround(g *rule.Ground, diagrams ...*statemachine.StateDiagram) *yongol.Fullstack {
	fs := &yongol.Fullstack{StateDiagrams: diagrams}
	fs.SetGround(g)
	return fs
}
