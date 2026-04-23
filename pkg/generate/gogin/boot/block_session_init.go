//ff:func feature=gen-gogin type=generator control=sequence
//ff:what blockSessionInit — session.Init (postgres 또는 memory) 블록

package boot

import "github.com/park-jun-woo/yongol/pkg/generate/prepared"

// blockSessionInit produces session initialization from a resolved
// Session. Callers guard with state.ActiveBackends.Session != nil so
// this function never sees an inactive subsystem — no raw manifest
// deref possible by signature.
func blockSessionInit(s prepared.Session) MainBlock {
	backend := s.Backend
	var lines []string
	if backend == "postgres" {
		lines = []string{
			`slog.Info("initializing session (postgres)")`,
			`sm, err := session.NewPostgresSession(ctx, conn)`,
			`if err != nil {`,
			`	slog.Error("session init", "err", err)`,
			`	os.Exit(1)`,
			`}`,
			`session.Init(sm)`,
		}
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
		Imports: []string{`"github.com/park-jun-woo/ssac/pkg/session"`},
		Lines:   lines,
	}
}
