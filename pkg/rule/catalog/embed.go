//ff:func feature=rule type=loader control=sequence topic=catalog
//ff:what Load — //go:embed 로 내장된 rulebook.md 를 파싱해 Catalog 반환 (failure → panic, 빌드 시점 감지)
package catalog

import (
	"bytes"
	_ "embed"
	"fmt"
	"log"
	"sync"
)

// rulebookSource is the canonical rulebook.md shipped with the yongol
// binary. Keep this file in sync with repo root rulebook.md — see the
// top-level note and the commit-time copy rule in CLAUDE.md.
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

// MustLoad returns the embedded catalog or calls log.Fatal on parse error.
// Intended for CLI entry points where recovery is not meaningful.
func MustLoad() *Catalog {
	c, err := Load()
	if err != nil {
		log.Fatalf("yongol: %v", err)
	}
	return c
}

// Source exposes the raw embedded rulebook bytes. Tests use this to
// re-parse the canonical source directly.
func Source() []byte {
	return rulebookSource
}
