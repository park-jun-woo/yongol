//ff:func feature=rule type=loader control=sequence topic=catalog
//ff:what Load — //go:embed 로 내장된 rulebook.md 를 파싱해 Catalog 반환 (sync.Once 캐시)
package catalog

import (
	"bytes"
	_ "embed"
	"fmt"
	"sync"
)

// rulebookSource is the canonical rulebook.md shipped with the yongol
// binary. It is a verbatim copy of the repo root rulebook.md, refreshed via
// `go generate ./pkg/rule/catalog` (see generate.go). Drift is caught by
// TestEmbedInSyncWithRootRulebook, which byte-compares the copy against the
// root file on every `go test`.
//
//go:embed rulebook.md
var rulebookSource []byte

var (
	loadOnce sync.Once
	cached   *Catalog
	cacheErr error
)

// Load returns the Catalog parsed from the embedded rulebook.md.
// The result is cached; subsequent calls return the same instance.
//
// Parsing failure is a build-time invariant violation (the rulebook is
// embedded) and surfaces as an error. Callers that prefer fatal semantics
// should use MustLoad.
func Load() (*Catalog, error) {
	loadOnce.Do(func() {
		rules, err := Parse(bytes.NewReader(rulebookSource))
		if err != nil {
			cacheErr = fmt.Errorf("parse embedded rulebook.md: %w", err)
			return
		}
		cached = NewCatalog(rules)
	})
	return cached, cacheErr
}
