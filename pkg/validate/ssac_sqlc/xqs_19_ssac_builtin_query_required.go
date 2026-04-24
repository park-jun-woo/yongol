//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-sqlc
//ff:what XQS-19 — SSaC 가 호출하는 DB-using ssac built-in 에 대응 sqlc 쿼리 존재 강제

package ssac_sqlc

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/ssacmeta"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xqs19SsacBuiltinQueryRequired validates XQS-19: when a SSaC function
// invokes a DB-using ssac built-in (`@call cache.Set`, `@publish topic`,
// `@subscribe topic`), every sqlc query declared as `used_by:` for that
// method in the package's interface.yaml must exist in fs.SQLcQueries.
//
// Catalog source: fs.SsacInterfaces — populated by parseSsacInterfaces
// (Phase002). The rule is catalog-free: adding a new `used_by:` entry to
// any interface.yaml port automatically brings the matching SSaC method
// under XQS-19 without yongol changes.
//
// Detection algorithm (per SSaC function):
//
//   for each sequence seq:
//     if seq is @call and interface.yaml[seq.Package] exists:
//       for each port whose UsedBy contains seq.Method:
//         require port.Name in fs.SQLcQueries
//     if seq is @publish:
//       resolve required ports against interface.yaml["queue"] where
//       UsedBy contains "Publish" (or "PublishTx")
//   if f.Subscribe != nil:
//     resolve ports for @subscribe against interface.yaml["queue"] where
//     UsedBy contains "Subscribe"
//
// Missing ports accumulate into one diagnostic per (func, method, query).
func xqs19SsacBuiltinQueryRequired(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil || len(fs.SsacInterfaces) == 0 {
		return nil
	}
	have := make(map[string]bool, len(fs.SQLcQueries))
	for _, q := range fs.SQLcQueries {
		have[q.Name] = true
	}

	var diags []diagnostic.Diagnostic
	for _, f := range fs.ServiceFuncs {
		for _, seq := range f.Sequences {
			pkg, method := resolveBuiltinCall(seq, f.Subscribe != nil)
			if pkg == "" {
				continue
			}
			iface := fs.SsacInterfaces[pkg]
			if iface == nil {
				continue
			}
			for _, port := range iface.Ports {
				if !containsUsedBy(port.UsedBy, method) {
					continue
				}
				if have[port.Name] {
					continue
				}
				diags = append(diags, buildXqs19Diag(f, seq, pkg, method, port.Name))
			}
		}
		if f.Subscribe != nil {
			iface := fs.SsacInterfaces["queue"]
			if iface == nil {
				continue
			}
			for _, port := range iface.Ports {
				if !containsUsedBy(port.UsedBy, "Subscribe") {
					continue
				}
				if have[port.Name] {
					continue
				}
				diags = append(diags, buildXqs19DiagSubscribe(f, port.Name))
			}
		}
	}
	return diags
}

// resolveBuiltinCall extracts the (pkg, method) pair that XQS-19 should
// check for a single SSaC sequence. Returns empty strings when the
// sequence is not a DB-facing built-in call.
//
//   @call <pkg>.<Method>   → (pkg, Method)
//   @publish "topic"       → ("queue", "Publish")
//
// @subscribe is handled one level up in the loop because it lives on the
// ServiceFunc, not the sequence.
func resolveBuiltinCall(seq ssacparser.Sequence, _ bool) (string, string) {
	switch seq.Type {
	case "call":
		if seq.Package == "" || seq.Model == "" {
			return "", ""
		}
		return seq.Package, seq.Model
	case "publish":
		return "queue", "Publish"
	}
	return "", ""
}

func containsUsedBy(usedBy []string, method string) bool {
	for _, m := range usedBy {
		if m == method {
			return true
		}
	}
	return false
}

func buildXqs19Diag(f ssacparser.ServiceFunc, seq ssacparser.Sequence, pkg, method, query string) diagnostic.Diagnostic {
	return diagnostic.Diagnostic{
		File:  f.FileName,
		Line:  seq.Line,
		Phase: diagnostic.PhaseValidate,
		Level: diagnostic.LevelError,
		Message: fmt.Sprintf(
			"[XQS-19] %s: @%s %s.%s requires sqlc query %q",
			f.Name, sequenceTag(seq.Type), pkg, method, query),
		Advice: buildXqs19Advice(pkg, query),
	}
}

func buildXqs19DiagSubscribe(f ssacparser.ServiceFunc, query string) diagnostic.Diagnostic {
	return diagnostic.Diagnostic{
		File:  f.FileName,
		Phase: diagnostic.PhaseValidate,
		Level: diagnostic.LevelError,
		Message: fmt.Sprintf(
			"[XQS-19] %s: @subscribe requires sqlc query %q",
			f.Name, query),
		Advice: buildXqs19Advice("queue", query),
	}
}

// sequenceTag renders the user-visible SSaC tag for a sequence Type so
// diagnostic messages match the authoring syntax (call / publish / …).
func sequenceTag(t string) string {
	return t
}

// buildXqs19Advice composes the copy-paste advice. interface.yaml's
// canonical_queries is the reference source; the package name is used to
// direct the user to specs/db/queries/<pkg>.sql.
func buildXqs19Advice(pkg, query string) string {
	return fmt.Sprintf(
		"Add a sqlc query named %q to specs/db/queries/%s.sql. "+
			"Refer to ssac/pkg/%s/interface.yaml canonical_queries for the template.",
		query, pkg, pkg)
}

// _ pins a reference to ssacmeta.Port so compilers flag an interface.yaml
// schema change (e.g. UsedBy rename) at the earliest possible test run.
var _ = ssacmeta.Port{}
