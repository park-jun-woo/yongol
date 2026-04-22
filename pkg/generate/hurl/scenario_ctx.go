//ff:type feature=gen-hurl type=model
//ff:what scenarioCtx — 시나리오 전체 빌드 중 공유되는 상태 (captures, roleMap)
package hurl

import (
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// scenarioCtx threads shared state through all phase builders.
//
// captures tracks hurl variables already produced by prior steps
// (e.g. "token", "token_admin", "workflow_id"). roleMap maps operationID
// to its required role from OPA policies; used by resolveTokenVar.
type scenarioCtx struct {
	fs       *yongol.Fullstack
	captures map[string]bool
	roleMap  map[string]string
}
