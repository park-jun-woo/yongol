//ff:func feature=validate type=rule control=sequence topic=manifest-infra
//ff:what XNQ-90 — manifest.queue.backend=postgres 시 canonical DDL + sqlc 쿼리 존재 강제

package manifest

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xnq90QueueBackendRequiresSQLC mirrors XNC-90 for the queue backend.
// Required entities (from ssac/pkg/queue/interface.yaml):
//
//   - DDL table: fullend_queue
//   - sqlc queries: QueuePublish / QueuePoll / QueueAck
//
// Note: QueuePoll / QueueAck are required by the interface.yaml even
// though Phase002's yongol wrapper currently only implements Publish /
// PublishTx. The validate rule enforces the full declaration so users
// who later wire a poll loop do not hit the missing-query at that time.
func xnq90QueueBackendRequiresSQLC(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	return validateBuiltinBackend(fs, backendSpec{
		Pkg:        "queue",
		Cfg:        queueCfg(fs),
		RequireDDL: "fullend_queue",
		RequireQueries: []string{
			"QueuePublish", "QueuePoll", "QueueAck",
		},
		RuleID: "XNQ-90",
	})
}

func queueCfg(fs *yongol.Fullstack) builtinBackend {
	if fs == nil || fs.Manifest == nil || fs.Manifest.Queue == nil {
		return builtinBackend{}
	}
	return builtinBackend{Present: true, Backend: fs.Manifest.Queue.Backend}
}
