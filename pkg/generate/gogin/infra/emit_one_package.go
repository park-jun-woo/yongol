//ff:func feature=gen-gogin type=generator control=sequence
//ff:what emitOnePackage — dispatch one ssac package to its registered emitter (or skip)

package infra

import (
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// emitOnePackage resolves the interface.yaml for pkg, filters to active
// ports, and dispatches to the matching packageEmitter. Inactive packages,
// dynamic-port packages (no static ports) and unregistered packages are
// silently skipped.
func emitOnePackage(fs *yongol.Fullstack, pkg string, mctx map[string]any, modulePath, artifactsDir string) error {
	iface := fs.SsacInterfaces[pkg]
	if iface == nil {
		return nil
	}
	// Dynamic-port packages (authz) are handled by Phase003 owner-lookup
	// emitter, not here — no static ports to walk.
	if len(iface.Ports) == 0 {
		return nil
	}
	active := activePorts(iface.Ports, mctx)
	if len(active) == 0 {
		return nil
	}
	emitter, ok := packageEmitters[pkg]
	if !ok {
		// No dedicated emitter registered — skip. A future package with
		// a pure sqlc-forwarder interface could be wired to
		// emitPostgresImplGeneric here.
		return nil
	}
	return emitter(iface, active, modulePath, artifactsDir)
}
