//ff:func feature=validate type=rule control=sequence topic=manifest-infra
//ff:what validateBuiltinBackend — XNC/XNS/XNQ-90 공용 검증 엔진

package manifest

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

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
	haveDDL := collectDDLNames(fs)
	haveQuery := collectQueryNames(fs)
	missing := missingBuiltinEntities(spec, haveDDL, haveQuery)
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
