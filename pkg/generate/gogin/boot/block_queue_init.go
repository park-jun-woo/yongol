//ff:func feature=gen-gogin type=generator control=iteration dimension=1
//ff:what blockQueueInit — queue.Init + Subscribe + Start + defer Close 블록

package boot

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// blockQueueInit produces queue initialization, subscriber registration, and
// Start/Close lifecycle. Active when manifest.queue.backend is set.
func blockQueueInit(fs *yongol.Fullstack) MainBlock {
	backend := "postgres"
	if fs.Manifest != nil && fs.Manifest.Queue != nil {
		backend = fs.Manifest.Queue.Backend
	}
	lines := []string{
		`slog.Info("initializing queue")`,
		fmt.Sprintf(`if err := queue.Init(ctx, %q, conn); err != nil {`, backend),
		`	slog.Error("queue init", "err", err)`,
		`	os.Exit(1)`,
		`}`,
		`defer queue.Close()`,
	}
	for _, fn := range fs.ServiceFuncs {
		if fn.Subscribe != nil {
			lines = append(lines,
				fmt.Sprintf(`queue.Subscribe(%q, srv.%s)`, fn.Subscribe.Topic, fn.Name),
			)
		}
	}
	lines = append(lines,
		`go func() {`,
		`	if err := queue.Start(ctx); err != nil {`,
		`		slog.Error("queue start", "err", err)`,
		`	}`,
		`}()`,
	)
	return MainBlock{
		Name:    "queue-init",
		Active:  hasQueue,
		Imports: []string{`"github.com/park-jun-woo/ssac/pkg/queue"`},
		Lines:   lines,
	}
}
