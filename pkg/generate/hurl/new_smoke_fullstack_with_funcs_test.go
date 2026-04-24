//ff:func feature=gen-hurl type=test-helper control=sequence
//ff:what newSmokeFullstackWithFuncs — explicit ServiceFuncs 로 Fullstack 조립

package hurl

import (
	"github.com/getkin/kin-openapi/openapi3"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// newSmokeFullstackWithFuncs is the explicit variant — caller provides
// the SSaC ServiceFuncs slice. Used by detect_auth_ops tests.
func newSmokeFullstackWithFuncs(doc *openapi3.T, funcs []ssac.ServiceFunc, diagrams ...*statemachine.StateDiagram) *yongol.Fullstack {
	return &yongol.Fullstack{
		Manifest:      &pmanifest.ProjectConfig{Backend: pmanifest.Backend{Auth: &pmanifest.Auth{Type: "jwt"}}},
		OpenAPIDoc:    doc,
		ServiceFuncs:  funcs,
		StateDiagrams: diagrams,
	}
}
