//ff:type feature=gen-hurl type=model
//ff:what scenarioCtx — 시나리오 전체 빌드 중 공유되는 상태 (captures, roleMap, auth detection)
package hurl

import (
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// scenarioCtx threads shared state through all phase builders.
//
// captures tracks hurl variables already produced by prior steps
// (e.g. "token", "token_admin", "workflow_id"). roleMap maps operationID
// to its required role from OPA policies; used by resolveTokenVar.
//
// authOpIDs / authSignup / authLogin are populated once per ctx by
// detectAuthOps (SSaC shape-based detection — BUG-023 fix). Downstream
// phases (create/read/update/delete step builders) call isAuthOpID(ctx,
// opID) to exclude auth endpoints from CRUD emission.
type scenarioCtx struct {
	fs       *yongol.Fullstack
	captures map[string]bool
	roleMap  map[string]string

	authOpIDs  map[string]authRole // opID → signup | login (empty map when no auth detected)
	authSignup *detectedAuthOp     // nil when no signup shape detected
	authLogin  *detectedAuthOp     // nil when no login shape detected
}
