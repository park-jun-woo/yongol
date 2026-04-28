//ff:func feature=gen-gogin type=generator control=sequence topic=auth-refresh
//ff:what emitAuthWrapper — ssac/pkg/auth.RefreshStore 를 사용자 sqlc Queries 로 구현하는 adapter 를 emit

package infra

import (
	"bytes"
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/ssacmeta"
)

// emitAuthWrapper writes `arts/backend/internal/infra/auth/postgres.go` — a
// postgres adapter that satisfies ssac/pkg/auth.RefreshStore by forwarding
// each method onto the user's sqlc Queries declared in
// ssac/pkg/auth/interface.yaml.
//
// Mapping (RefreshStore method → interface.yaml port):
//
//   Create       → RefreshTokenInsert
//   Consume      → RefreshTokenFindByHash + RefreshTokenRevoke
//                  (SELECT active row, validate, then UPDATE revoked_at —
//                   this is the only RefreshStore method that needs two
//                   queries; every other method is a single forwarder.)
//   Revoke       → RefreshTokenRevoke
//   RevokeAll    → RefreshTokenRevokeAll
//
// Type-translation glue:
//
//   - RefreshStore.Create takes `claims any`; the wrapper json.Marshals
//     it so the sqlc param stays []byte (stored as JSONB).
//   - RefreshStore.Consume plaintext token → sha256 hash before every
//     DB touch via auth.HashRefreshToken. Revoked rows surface as
//     ErrRefreshTokenReused together with their claims so reuse-detection
//     lockout can inspect the claim set.
//   - expires_at / revoked_at are pgtype.Timestamptz for pgx/v5 params;
//     the wrapper converts to/from time.Time at the boundary.
//
// LoginLookup is declared in interface.yaml `when: always` but is NOT
// implemented by this adapter — it is consumed directly by user SSaC
// @call auth.Login handlers (Phase003 handler codegen), not by the
// RefreshStore interface.
func emitAuthWrapper(iface *ssacmeta.PackageInterface, active []ssacmeta.Port, modulePath, artifactsDir string) error {
	insertPort := portByName(active, "RefreshTokenInsert")
	findPort := portByName(active, "RefreshTokenFindByHash")
	revokePort := portByName(active, "RefreshTokenRevoke")
	revokeAllPort := portByName(active, "RefreshTokenRevokeAll")
	if insertPort == nil || findPort == nil || revokePort == nil || revokeAllPort == nil {
		return fmt.Errorf("auth: interface.yaml missing one of RefreshTokenInsert/FindByHash/Revoke/RevokeAll (active ports: %d)", len(active))
	}

	var buf bytes.Buffer
	fmt.Fprintf(&buf, authWrapperTemplate, modulePath, insertPort.Name, findPort.Name, revokePort.Name, revokeAllPort.Name)

	return writeAdapterFile(artifactsDir, iface.Package, buf.Bytes())
}
