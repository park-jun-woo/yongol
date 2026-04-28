//ff:func feature=gen-gogin type=util control=sequence
//ff:what buildAuth — @auth 시퀀스 빌더 (OwnerLookup<Resource> + Owners 맵 + authz.Check)

package ssac

import (
	"fmt"
	"strings"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// buildAuth emits the authz.Check call for an `@auth` sequence. The entire
// *model.UserClaim struct is forwarded as the `Claim` field; rego.Input
// json-marshals it using the JSON tags produced by generate_user_claim.go
// into the `input.claims` object expected by the policy.
func (g *methodGen) buildAuth(seq ssacparser.Sequence) ([]string, []string) {
	status := resolveErrStatus("auth", seq.ErrStatus)
	msg := seq.Message
	if msg == "" {
		msg = neutralMessage(status)
	}
	extraFields := g.mapFields(seq.Inputs)

	// `assign` selects := vs = for the final authz.Check call. When the
	// ownership branch below runs it re-evaluates assignOp after emitting
	// the owner lookup (which itself declares err via :=), so `assign`
	// must remain reactive to methodGen's FirstErr state. (BUG-029)
	assign := g.assignOp(false)

	mapping := findOwnershipMapping(g.Ownerships, seq.Resource)
	imports := []string{`"github.com/park-jun-woo/ssac/pkg/authz"`, `"log/slog"`}

	ownerLines, ownersExpr, ownerImports := g.ownershipLookup(seq, mapping, status, msg)
	imports = append(imports, ownerImports...)
	if len(ownerLines) > 0 {
		// After emitting the initial lookup, subsequent err assignments in
		// this block must reuse `=` (the first := already declared err).
		assign = g.assignOp(false)
	}

	checkFields := []string{
		"Ctx: ctx",
		fmt.Sprintf("Action: %q", seq.Action),
		fmt.Sprintf("Resource: %q", seq.Resource),
		"Claim: currentUser",
	}
	if extraFields != "" {
		checkFields = append(checkFields, extraFields)
	}
	checkFields = append(checkFields, "Owners: "+ownersExpr)

	lines := ownerLines
	lines = append(lines,
		fmt.Sprintf("_, err %s authz.Check(authz.CheckRequest{%s})", assign, strings.Join(checkFields, ", ")),
		"if err != nil {",
		fmt.Sprintf("\t%s(\"handler: %s\", \"op\", %q, \"status\", %d, \"err\", err)", logLevelFuncForStatus(status), logTagForStatus(status), g.FuncName, status),
		fmt.Sprintf("\treturn api.%s%dJSONResponse{Error: %q, Code: strPtr(%q)}, nil", g.FuncName, status, msg, neutralCode(status)),
		"}",
	)
	return lines, imports
}
