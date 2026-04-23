//ff:func feature=gen-gogin type=generator control=sequence
//ff:what blockQueueInit — queue.Init + Subscribe + Start + defer Close 블록

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
func blockQueueInit(q prepared.Queue, serviceFuncs []ssac.ServiceFunc) MainBlock {
	lines := []string{
		`slog.Info("initializing queue")`,
		fmt.Sprintf(`if err := queue.Init(ctx, %q, conn); err != nil {`, q.Backend),
		`	slog.Error("queue init", "err", err)`,
		`	os.Exit(1)`,
		`}`,
		`defer queue.Close()`,
	}
	lines = appendQueueSubscribeLines(lines, serviceFuncs)
	lines = append(lines,
		`go func() {`,
		`	if err := queue.Start(ctx); err != nil {`,
		`		slog.Error("queue start", "err", err)`,
		`	}`,
		`}()`,
	)
	return MainBlock{
		Name: "queue-init",
		// Active left nil: collectActiveBlocks appends this block only
		// when prepared.State.ActiveBackends.Queue != nil.
		Imports: []string{`"github.com/park-jun-woo/ssac/pkg/queue"`},
		Lines:   lines,
	}
}
