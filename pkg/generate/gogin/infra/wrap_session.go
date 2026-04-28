//ff:func feature=gen-gogin type=generator control=sequence topic=session
//ff:what emitSessionWrapper — ssac/pkg/session.SessionModel 을 사용자 sqlc Queries 로 구현하는 adapter 를 emit

package infra

import (
	"bytes"
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/ssacmeta"
)

// emitSessionWrapper writes `arts/backend/internal/infra/session/postgres.go`.
// The shape mirrors wrap_cache — session and cache share the same
// TTL-keyed interface (Set/Get/Delete), the only differences being the
// fullend_sessions table, the Session* sqlc names, and the SessionModel
// interface. The same `any → JSON`, `[]byte → string`, `ttl → expires_at`
// glue applies.
func emitSessionWrapper(iface *ssacmeta.PackageInterface, active []ssacmeta.Port, modulePath, artifactsDir string) error {
	setPort := portByName(active, "SessionSet")
	getPort := portByName(active, "SessionGet")
	delPort := portByName(active, "SessionDelete")
	if setPort == nil || getPort == nil || delPort == nil {
		return fmt.Errorf("session: interface.yaml missing one of SessionSet/SessionGet/SessionDelete (active ports: %d)", len(active))
	}

	var buf bytes.Buffer
	fmt.Fprintf(&buf, sessionWrapperTemplate, modulePath, setPort.Name, getPort.Name, delPort.Name)

	return writeAdapterFile(artifactsDir, iface.Package, buf.Bytes())
}
