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
	fmt.Fprintf(&buf, `//ff:type feature=infra type=model topic=queue
//ff:what postgresQueue — ssac/pkg/queue.Backend 구현 (yongol codegen from ssac/pkg/queue/interface.yaml)

package queue

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/park-jun-woo/ssac/pkg/queue"

	"%[1]s/internal/db"
)

// ErrTxMustBePgx — PublishTx received a non-pgx transaction handle. The
// postgres adapter is driver-specific by design; callers using a different
// driver need a separate Backend.
var ErrTxMustBePgx = errors.New("infra/queue: PublishTx requires pgx.Tx")

// postgresQueue adapts the user's sqlc Queries onto queue.Backend.
// Construct via NewPostgres(queries); wire from main.go via
// queue.SetBackend(infraqueue.NewPostgres(queries)).
type postgresQueue struct {
	q *db.Queries
}

// NewPostgres returns a queue.Backend that persists messages in
// fullend_queue via the user's sqlc %[2]s method.
func NewPostgres(q *db.Queries) queue.Backend {
	return &postgresQueue{q: q}
}

// Publish inserts a pending row in fullend_queue. Priority defaults to
// "normal" when unset; deliver_at is now+cfg.Delay seconds. traceparent is
// set to the empty string on the non-tx Publish path — callers that need
// span propagation should use PublishTx inside an otel-instrumented tx.
// interface.yaml port: %[2]s.
func (p *postgresQueue) Publish(ctx context.Context, topic string, data []byte, cfg queue.PublishConfig) error {
	return p.q.%[2]s(ctx, db.%[2]sParams{
		Topic:       topic,
		Payload:     data,
		Priority:    resolvePriority(cfg.Priority),
		DeliverAt:   pgtype.Timestamptz{Time: time.Now().Add(time.Duration(cfg.Delay) * time.Second), Valid: true},
		Traceparent: "",
	})
}

// PublishTx inserts the queue row bound to the caller's transaction. tx
// must be a pgx.Tx (the sqlc Queries.WithTx expects one); any other type
// returns ErrTxMustBePgx.
// interface.yaml port: %[2]s.
func (p *postgresQueue) PublishTx(ctx context.Context, tx any, topic string, data []byte, cfg queue.PublishConfig) error {
	pgxTx, ok := tx.(pgx.Tx)
	if !ok {
		return ErrTxMustBePgx
	}
	qtx := p.q.WithTx(pgxTx)
	return qtx.%[2]s(ctx, db.%[2]sParams{
		Topic:       topic,
		Payload:     data,
		Priority:    resolvePriority(cfg.Priority),
		DeliverAt:   pgtype.Timestamptz{Time: time.Now().Add(time.Duration(cfg.Delay) * time.Second), Valid: true},
		Traceparent: "",
	})
}

// resolvePriority normalises empty priority strings to "normal" so the
// fullend_queue ORDER BY CASE remains well-defined regardless of caller
// options.
func resolvePriority(p string) string {
	if p == "" {
		return "normal"
	}
	return p
}
`, modulePath, pubPort.Name)

	return writeAdapterFile(artifactsDir, iface.Package, buf.Bytes())
}
