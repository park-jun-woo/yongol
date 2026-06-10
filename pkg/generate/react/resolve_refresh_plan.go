//ff:func feature=gen-react type=util control=sequence
//ff:what resolveRefreshPlan — refresh_op 선언 우선, 미선언 시 구조 추론 (후보 2+ → generate ERROR)

package react

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// resolveRefreshPlan decides the bearer 401 refresh flow. nil (no error) is
// the explicit downgrade — Bearer injection only, 401 clears + /login:
//   - frontend.auth or its refresh_field is undeclared, or
//   - no usable refresh operation exists (declared op missing is XON-60's
//     diagnosis at validate time; zero inference candidates is a legitimate
//     "no refresh endpoint" project).
//
// A declared frontend.auth.refresh_op wins. Without it the op is inferred
// structurally: operations whose 2xx response carries both token_field and
// refresh_field, minus ops bound by an STML data-capture declaration (those
// are login-style token producers, e.g. Login). Two or more candidates is a
// generate ERROR asking for an explicit refresh_op declaration. A declared
// op that takes path parameters is also an ERROR — the generated 401 hook
// has no way to supply them.
func resolveRefreshPlan(fs *yongol.Fullstack) (*refreshPlan, error) {
	fa := fs.Manifest.Frontend.Auth
	doc := fs.OpenAPIDoc
	if fa == nil || fa.RefreshField == "" || doc == nil || doc.Paths == nil {
		return nil, nil
	}
	if fa.RefreshOp != "" {
		rp, found := findRefreshOp(doc, fa)
		if !found {
			return nil, nil // XON-60 reports the missing operationId
		}
		if strings.Contains(rp.path, "{") {
			return nil, fmt.Errorf("frontend.auth.refresh_op %q takes path parameters (%s %s) — the generated 401 refresh call cannot supply them; declare a parameterless refresh operation", fa.RefreshOp, rp.method, rp.path)
		}
		return rp, nil
	}
	candidates := inferRefreshOps(doc, fa, captureDeclaredOps(fs.STMLPages))
	if len(candidates) == 0 {
		return nil, nil
	}
	if len(candidates) > 1 {
		var ids []string
		for _, c := range candidates {
			ids = append(ids, c.opID)
		}
		return nil, fmt.Errorf("ambiguous refresh op: %d operations carry both %q and %q in a 2xx response without an STML data-capture declaration (%s); declare frontend.auth.refresh_op in manifest.yaml to pick one", len(candidates), fa.TokenField, fa.RefreshField, strings.Join(ids, ", "))
	}
	return &candidates[0], nil
}
