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

	// Separate ResourceID from the rest of Inputs — it needs special
	// stringification for authz.CheckRequest.ResourceID (string).
	filtered := make(map[string]string, len(seq.Inputs))
	for k, v := range seq.Inputs {
		if k != "ResourceID" {
			filtered[k] = v
		}
	}
	extraFields := g.mapFields(filtered)

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

	// Emit ResourceID with explicit string conversion based on PK type.
	rawRID, hasRID := seq.Inputs["ResourceID"]
	if hasRID && !isResourceIDZero(rawRID) {
		ridExpr := g.mapValue(rawRID)
		pkCol := g.lookupResourcePKColumn(seq.Resource)
		if pkCol != nil && isUUIDColumn(pkCol) {
			checkFields = append(checkFields, "ResourceID: pgtypex.UUIDToString(pgtypex.ToPgUUID("+ridExpr+"))")
			imports = append(imports, `"github.com/park-jun-woo/ssac/pkg/pgtypex"`)
		} else {
			checkFields = append(checkFields, "ResourceID: strconv.FormatInt("+ridExpr+", 10)")
			imports = append(imports, `"strconv"`)
		}
	}

	checkFields = append(checkFields, "Owners: "+ownersExpr)

	lines := ownerLines
	lines = append(lines,
		fmt.Sprintf("_, err %s authz.Check(authz.CheckRequest{%s})", assign, strings.Join(checkFields, ", ")),
		"if err != nil {",
		fmt.Sprintf("\t%s(\"handler: %s\", \"op\", %q, \"status\", %d, \"err\", err)", logLevelFuncForStatus(status), logTagForStatus(status), g.FuncName, status),
		"\t" + g.guardReturn(msg, status),
		"}",
	)
	if g.IsSubscribe {
		imports = append(imports, `"fmt"`)
	}
	return lines, imports
}
