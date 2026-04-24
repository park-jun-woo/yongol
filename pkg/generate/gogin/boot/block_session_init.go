//ff:func feature=gen-gogin type=generator control=sequence
//ff:what blockSessionInit — session.Init (postgres infra 어댑터 또는 memory) 블록

package boot

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
)

// blockSessionInit produces session initialization from a resolved
// Session. Callers guard with state.ActiveBackends.Session != nil so
// this function never sees an inactive subsystem — no raw manifest
// deref possible by signature.
//
// Phase002 (ssac/purify) — postgres branch now instantiates the
// yongol-generated adapter at `<module>/internal/infra/session` via
// session.NewPostgres(queries). The ssac-side session.NewPostgresSession
// constructor was removed in Phase001.
func blockSessionInit(s prepared.Session, modulePath string) MainBlock {
	backend := s.Backend
	var lines []string
	imports := []string{
		`"github.com/park-jun-woo/ssac/pkg/session"`,
	}
	if backend == "postgres" {
		lines = []string{
			`slog.Info("initializing session (postgres)")`,
			`session.Init(infrasession.NewPostgres(queries))`,
		}
		imports = append(imports, fmt.Sprintf(`infrasession "%s/internal/infra/session"`, modulePath))
	} else {
		lines = []string{
			`slog.Info("initializing session (memory)")`,
			`session.Init(session.NewMemorySession())`,
		}
	}
	return MainBlock{
		Name: "session-init",
		// Active left nil: collectActiveBlocks appends this block only
		// when prepared.State.ActiveBackends.Session != nil, so the
		// gate moved out of Active into the caller.
		Imports: imports,
		Lines:   lines,
	}
}
