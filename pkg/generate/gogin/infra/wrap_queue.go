//ff:func feature=gen-gogin type=generator control=sequence topic=queue
//ff:what emitQueueWrapper — ssac/pkg/queue.Backend 를 사용자 sqlc Queries 로 구현하는 adapter 를 emit

package infra

import (
	"bytes"
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/ssacmeta"
)

// emitQueueWrapper writes `arts/backend/internal/infra/queue/postgres.go` —
// a postgres adapter that satisfies ssac/pkg/queue.Backend by forwarding
// Publish/PublishTx onto the user's sqlc-generated QueuePublish method
// (declared in ssac/pkg/queue/interface.yaml).
//
// Type-translation glue:
//
//   - Backend.Publish takes `cfg queue.PublishConfig`; sqlc QueuePublish
//     expects concrete (topic, payload, priority, deliver_at, traceparent).
//     The wrapper maps Priority/Delay onto those columns and extracts
//     traceparent from the ctx via queue.TraceparentFromContext (W3C
//     TraceContext propagation added in ssac queue Phase005).
//
//   - Backend.PublishTx takes `tx any` so ssac does not depend on pgx.
//     The wrapper asserts `pgx.Tx` and forwards via `q.WithTx(tx).QueuePublish`.
//     A non-pgx tx falls through to an error — memory semantics are
//     preserved by the memoryBackend inside ssac, not here.
//
// Scope (Phase002):
//
//   - QueuePublish is fully implemented.
//   - QueuePoll / QueueAck (declared in interface.yaml) require an
//     outbox polling loop (queue.Starter). That loop needs access to
//     ssac's package-private handlers map, so implementing it requires
//     a ssac-side export first. Tracked separately; zenflow smoke does
//     not exercise durable queues.
func emitQueueWrapper(iface *ssacmeta.PackageInterface, active []ssacmeta.Port, modulePath, artifactsDir string) error {
	pubPort := portByName(active, "QueuePublish")
	if pubPort == nil {
		return fmt.Errorf("queue: interface.yaml missing QueuePublish (active ports: %d)", len(active))
	}

	var buf bytes.Buffer
	fmt.Fprintf(&buf, queueWrapperTemplate, modulePath, pubPort.Name)

	return writeAdapterFile(artifactsDir, iface.Package, buf.Bytes())
}
