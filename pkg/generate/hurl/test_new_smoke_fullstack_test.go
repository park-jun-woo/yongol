//ff:func feature=gen-hurl type=test-helper control=sequence
//ff:what newSmokeFullstack — OpenAPI doc + state diagram 을 smoke 테스트용 Fullstack 에 조립

package hurl

import (
	"github.com/getkin/kin-openapi/openapi3"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// newSmokeFullstack assembles a minimal *yongol.Fullstack wired with the
// supplied OpenAPI doc and (optionally) state diagrams. Manifest.Auth is
// populated so buildAuthSteps fires.
func newSmokeFullstack(doc *openapi3.T, diagrams ...*statemachine.StateDiagram) *yongol.Fullstack {
	return &yongol.Fullstack{
		Manifest:      &pmanifest.ProjectConfig{Backend: pmanifest.Backend{Auth: &pmanifest.Auth{Type: "jwt"}}},
		OpenAPIDoc:    doc,
		StateDiagrams: diagrams,
	}
}
