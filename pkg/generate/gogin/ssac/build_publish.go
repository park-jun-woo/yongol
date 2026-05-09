//ff:func feature=gen-gogin type=util control=sequence
//ff:what buildPublish — @publish 시퀀스 빌더 (tx 내부: queue.PublishTx / 그 외: queue.Publish)

package ssac

import (
	"fmt"
	"sort"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// buildPublish emits the code for a @publish sequence.
//
// Phase006 (outbox atomicity): when the enclosing handler runs inside a DB
// transaction (g.UseTx == true), the enqueue INSERT must share that tx so a
// failed commit also drops the event. We emit queue.PublishTx(ctx, tx, ...)
// and propagate the error — the surrounding defer rollback then voids the
// whole operation.
//
// For tx-less handlers (e.g. pure read paths that still fan out an event) we
// keep the legacy best-effort queue.Publish call and only log failures, since
// there is no business transaction to unwind.
func (g *methodGen) buildPublish(seq ssacparser.Sequence) ([]string, []string) {
	keys := make([]string, 0, len(seq.Inputs))
	for k := range seq.Inputs {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var fields []string
	for _, k := range keys {
		fields = append(fields, fmt.Sprintf("\t%q: %s,", k, g.mapValue(seq.Inputs[k])))
	}

	imports := []string{`"github.com/park-jun-woo/ssac/pkg/queue"`}

	if g.UseTx {
		return buildPublishTx(seq, fields, imports), imports
	}
	return g.buildPublishBestEffort(seq, fields, &imports), imports
}
