//ff:func feature=gen-hurl type=command control=sequence
//ff:what Generate — OpenAPI+SSaC+StateDiagram+Policy+DDL 조합으로 smoke.hurl 자동 생성
package hurl

import (
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// Generate produces arts/tests/smoke.hurl from a parsed Fullstack and
// mirrors user-authored scenario-*.hurl / invariant-*.hurl from
// specs/tests/ into arts/tests/ so they are executable from the same
// location as smoke.hurl (BUG-026).
func Generate(fs *yongol.Fullstack, artifactsDir string) error {
	steps := buildScenarioOrder(fs)
	if err := writeSmokeHurl(steps, artifactsDir); err != nil {
		return err
	}
	if fs != nil {
		if err := mirrorUserHurlFiles(fs.SpecsDir, artifactsDir); err != nil {
			return err
		}
	}
	return nil
}
