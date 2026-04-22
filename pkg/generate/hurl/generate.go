//ff:func feature=gen-hurl type=command control=sequence
//ff:what Generate — OpenAPI+SSaC+StateDiagram+Policy+DDL 조합으로 smoke.hurl 자동 생성
package hurl

import (
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// Generate produces arts/tests/smoke.hurl from a parsed Fullstack.
func Generate(fs *yongol.Fullstack, artifactsDir string) error {
	steps := buildScenarioOrder(fs)
	return writeSmokeHurl(steps, artifactsDir)
}
