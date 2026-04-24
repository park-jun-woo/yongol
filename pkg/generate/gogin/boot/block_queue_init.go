//ff:func feature=gen-gogin type=generator control=sequence
//ff:what blockQueueInit — queue.Init + SetBackend (memory or postgres infra 어댑터) + Subscribe + Start + defer Close

package boot

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// blockQueueInit produces queue initialization, subscriber registration,
// and Start/Close lifecycle. Caller guards with
// state.ActiveBackends.Queue != nil; q.Backend is the resolved default
// (matches pre-Phase001 behavior: "postgres" when manifest silent).
// serviceFuncs carries the Subscribe metadata so this function stays
// free of raw fs access.
//
// Phase002 (ssac/purify) — durable (postgres) backends no longer accept
// ssac-native Init("postgres", conn). ssac accepts only "memory" as a
// built-in backend; the postgres driver is provided by yongol-generated
// infra at `<module>/internal/infra/queue`. The emitter wires the
// adapter through queue.SetBackend after Init("memory").
func blockQueueInit(q prepared.Queue, serviceFuncs []ssac.ServiceFunc, modulePath string) MainBlock {
	lines := []string{
		`slog.Info("initializing queue")`,
		`if err := queue.Init(ctx, "memory"); err != nil {`,
		`	slog.Error("queue init", "err", err)`,
		`	os.Exit(1)`,
		`}`,
	}
	imports := []string{`"github.com/park-jun-woo/ssac/pkg/queue"`}
	if q.Backend == "postgres" {
		lines = append(lines,
			`queue.SetBackend(infraqueue.NewPostgres(queries))`,
		)
		imports = append(imports, fmt.Sprintf(`infraqueue "%s/internal/infra/queue"`, modulePath))
	}
	lines = append(lines, `defer queue.Close()`)
	lines = appendQueueSubscribeLines(lines, serviceFuncs)
	lines = append(lines,
		`go func() {`,
		`	if err := queue.Start(ctx); err != nil {`,
		`		slog.Error("queue start", "err", err)`,
		`	}`,
		`}()`,
	)
	_ = q // silence unused when backend != postgres — q.Backend is still read above
	return MainBlock{
		Name: "queue-init",
		// Active left nil: collectActiveBlocks appends this block only
		// when prepared.State.ActiveBackends.Queue != nil.
		Imports: imports,
		Lines:   lines,
	}
}
