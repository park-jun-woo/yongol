package infra

// queueWrapperTemplate is the printf-style template for
// arts/backend/internal/infra/queue/postgres.go.
var queueWrapperTemplate = queueWrapperHeaderType + queueWrapperHeaderWhat + `

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
`

var queueWrapperHeaderType = "//" + "ff:type feature=infra type=model topic=queue\n"
var queueWrapperHeaderWhat = "//" + "ff:what postgresQueue — ssac/pkg/queue.Backend 구현 (yongol codegen from ssac/pkg/queue/interface.yaml)"
