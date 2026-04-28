//ff:func feature=validate type=rule control=sequence topic=manifest-infra
//ff:what canonicalAdvice — ssac interface.yaml 의 canonical_ddl + canonical_queries 를 advice 로 렌더

package manifest

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

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
