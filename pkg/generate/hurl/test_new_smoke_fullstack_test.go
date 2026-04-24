//ff:func feature=gen-hurl type=test-helper control=sequence
//ff:what newSmokeFullstack — OpenAPI doc + state diagram + 자동 SSaC auth funcs 를 Fullstack 으로 조립

package hurl

import (
	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// newSmokeFullstack assembles a minimal *yongol.Fullstack wired with
// the supplied OpenAPI doc and (optionally) state diagrams. Manifest.
// Auth is populated so buildAuthSteps fires.
//
// To keep Phase003 shape detection honest without every test spelling
// out the SSaC stanza, any operationId whose path starts with
// "/auth/register" or "/auth/signup" gets a synthetic signup-shape
// ServiceFunc, and "/auth/login" / "/auth/signin" get a login-shape
// one. Tests can call newSmokeFullstackWithFuncs when explicit SSaC
// fixtures are needed.
func newSmokeFullstack(doc *openapi3.T, diagrams ...*statemachine.StateDiagram) *yongol.Fullstack {
	return newSmokeFullstackWithFuncs(doc, inferAuthServiceFuncs(doc), diagrams...)
}
