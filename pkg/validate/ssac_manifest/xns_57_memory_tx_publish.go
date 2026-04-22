//ff:func feature=validate type=rule control=sequence topic=config-check
//ff:what XNS-57 — warns when queue.backend=memory is combined with a tx-bound @publish

package ssac_manifest

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// xns57MemoryTxPublish warns when manifest queue.backend is "memory" and any
// SSaC service func mixes a mutating sequence (@post/@put/@delete) with a
// @publish — yongol generates queue.PublishTx(ctx, tx, ...) for that
// combination, but the memory backend cannot honour the tx semantics and
// returns ErrTxUnsupported at runtime, failing every such handler.
//
// Memory backends are acceptable for tests and tx-less @publish (e.g. read
// handlers emitting audit events). Production should use "postgres".
func xns57MemoryTxPublish(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil {
		return nil
	}
	backend := queueBackend(fs)
	if backend != "memory" {
		return nil
	}

	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		if !hasTxBoundPublish(fn) {
			continue
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:  fn.FileName,
			Phase: diagnostic.PhaseValidate,
			Level: diagnostic.LevelWarning,
			Message: fmt.Sprintf(
				"[XNS-57] queue.backend is \"memory\" but %s mixes @publish with @post/@put/@delete — tx-bound publish has no atomicity guarantee",
				fn.Name,
			),
			Advice: "In production, set queue.backend to \"postgres\" (the memory backend does not support PublishTx)",
		})
	}
	return diags
}

// queueBackend returns the manifest queue.backend value ("postgres"/"memory"/...)
// or "" if unset. Ground's Config map only flags key presence, so we read the
// manifest struct directly.
func queueBackend(fs *yongol.Fullstack) string {
	if fs.Manifest == nil || fs.Manifest.Queue == nil {
		return ""
	}
	return fs.Manifest.Queue.Backend
}

// hasTxBoundPublish returns true when the service func has both a mutating
// sequence (which triggers tx generation in yongol codegen) and a @publish
// sequence. This mirrors pkg/generate/gogin/ssac/needs_transaction.go.
func hasTxBoundPublish(fn ssacparser.ServiceFunc) bool {
	var hasMutation, hasPublish bool
	for _, seq := range fn.Sequences {
		switch seq.Type {
		case "post", "put", "delete":
			hasMutation = true
		case "publish":
			hasPublish = true
		}
	}
	return hasMutation && hasPublish
}
