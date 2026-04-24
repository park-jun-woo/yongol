//ff:func feature=validate type=rule control=sequence topic=manifest-infra
//ff:what XNC-90 — manifest.cache.backend=postgres 시 canonical DDL + sqlc 쿼리 존재 강제

package manifest

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xnc90CacheBackendRequiresSQLC validates XNC-90: when the user opts into
// manifest.cache.backend == "postgres", the companion DDL table
// `fullend_cache` and sqlc queries `CacheSet / CacheGet / CacheDelete`
// (declared as ports in ssac/pkg/cache/interface.yaml) must exist in the
// user-authored SSOTs. Missing entries surface as a single ERROR so the
// user sees every required entity in one advice blob.
//
// Advice payload: the corresponding ssac interface.yaml's canonical_ddl +
// canonical_queries blocks concatenated verbatim. This keeps the emitter
// catalog-free — renaming a port in ssac automatically propagates to the
// advice text the next time yongol is rebuilt.
func xnc90CacheBackendRequiresSQLC(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	return validateBuiltinBackend(fs, backendSpec{
		Pkg:        "cache",
		Cfg:        cacheCfg(fs),
		RequireDDL: "fullend_cache",
		RequireQueries: []string{
			"CacheSet", "CacheGet", "CacheDelete",
		},
		RuleID: "XNC-90",
	})
}

func cacheCfg(fs *yongol.Fullstack) builtinBackend {
	if fs == nil || fs.Manifest == nil || fs.Manifest.Cache == nil {
		return builtinBackend{}
	}
	return builtinBackend{Present: true, Backend: fs.Manifest.Cache.Backend}
}

// builtinBackend is a small carrier used by validateBuiltinBackend so the
// signature stays stable for the 4 symmetrical rules (cache / session /
// queue / auth.refresh).
type builtinBackend struct {
	Present bool
	Backend string
}

// backendSpec declares everything validateBuiltinBackend needs to check
// one backend. Defined as a struct rather than positional args so future
// additions (e.g. an optional second DDL table) don't break call sites.
type backendSpec struct {
	Pkg            string
	Cfg            builtinBackend
	RequireDDL     string
	RequireQueries []string
	RuleID         string
}

// validateBuiltinBackend is the shared engine for XNC-90 / XNS-90 / XNQ-90.
// It short-circuits when the backend is absent or memory-only, then checks
// that the DDL table and every sqlc query name exist. On any miss it emits
// a single diagnostic whose advice concatenates the interface.yaml's
// canonical_ddl + canonical_queries blocks.
func validateBuiltinBackend(fs *yongol.Fullstack, spec backendSpec) []diagnostic.Diagnostic {
	if fs == nil || !spec.Cfg.Present {
		return nil
	}
	if strings.ToLower(spec.Cfg.Backend) != "postgres" {
		return nil
	}

	haveDDL := make(map[string]bool, len(fs.DDLTables))
	for _, t := range fs.DDLTables {
		haveDDL[t.Name] = true
	}
	haveQuery := make(map[string]bool, len(fs.SQLcQueries))
	for _, q := range fs.SQLcQueries {
		haveQuery[q.Name] = true
	}

	var missing []string
	if spec.RequireDDL != "" && !haveDDL[spec.RequireDDL] {
		missing = append(missing, "DDL table "+spec.RequireDDL)
	}
	for _, qn := range spec.RequireQueries {
		if !haveQuery[qn] {
			missing = append(missing, "sqlc query "+qn)
		}
	}
	if len(missing) == 0 {
		return nil
	}

	file := "manifest.yaml"
	if fs.SpecsDir != "" {
		file = fs.SpecsDir + "/manifest.yaml"
	}
	return []diagnostic.Diagnostic{{
		File:    file,
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelError,
		Message: fmt.Sprintf("[%s] manifest.%s.backend=postgres but missing: %s", spec.RuleID, spec.Pkg, strings.Join(missing, ", ")),
		Advice:  canonicalAdvice(fs, spec.Pkg),
	}}
}

// canonicalAdvice renders the ssac interface.yaml's canonical_ddl +
// canonical_queries as a copy-paste-ready block. Falls back to a
// "interface.yaml not found" note when ssac sources are unavailable at
// validate time (e.g. running against a module-cache ssac that omits the
// yaml) — this keeps the rule useful even without a local ssac checkout.
func canonicalAdvice(fs *yongol.Fullstack, pkg string) string {
	if fs == nil || fs.SsacInterfaces == nil {
		return "Refer to ssac/pkg/" + pkg + "/interface.yaml for the canonical DDL + queries."
	}
	iface := fs.SsacInterfaces[pkg]
	if iface == nil {
		return "Refer to ssac/pkg/" + pkg + "/interface.yaml for the canonical DDL + queries."
	}
	var b strings.Builder
	b.WriteString("Add the canonical DDL + sqlc queries declared in ssac/pkg/")
	b.WriteString(pkg)
	b.WriteString("/interface.yaml:\n\n")
	if iface.CanonicalDDL != "" {
		b.WriteString("-- specs/db/" + pkg + ".sql --\n")
		b.WriteString(strings.TrimRight(iface.CanonicalDDL, "\n"))
		b.WriteString("\n\n")
	}
	if iface.CanonicalQueries != "" {
		b.WriteString("-- specs/db/queries/" + pkg + ".sql --\n")
		b.WriteString(strings.TrimRight(iface.CanonicalQueries, "\n"))
		b.WriteString("\n")
	}
	return b.String()
}
