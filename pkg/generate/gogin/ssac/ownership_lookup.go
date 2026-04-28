//ff:func feature=gen-gogin type=util control=sequence
//ff:what ownershipLookup — emit OwnerLookup<Res> call + Owners literal for buildAuth

package ssac

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/parser/rego"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// ownershipLookup emits the `<owner>, err := qtx.OwnerLookup<Res>(ctx, rid)`
// preamble and builds the Owners map literal used by authz.CheckRequest.
// Returns empty lines + "nil" owners expression when no ownership mapping
// applies (no matching @ownership rule or the resource-id is statically
// zero — see BUG-033).
func (g *methodGen) ownershipLookup(seq ssacparser.Sequence, mapping *rego.OwnershipMapping, status int, msg string) ([]string, string, []string) {
	if mapping == nil {
		return nil, "nil", nil
	}
	rawRID, hasRID := seq.Inputs["ResourceID"]
	if !hasRID || isResourceIDZero(rawRID) {
		return nil, "nil", nil
	}
	ridExpr := g.mapValue(rawRID)
	if ridExpr == "" {
		ridExpr = "0"
	}
	queryName := ownerLookupQueryName(mapping.Resource)
	ownerVar := "owner" + pascalCase(mapping.Resource)
	queriesRecv := "qtx"
	if !g.UseTx {
		queriesRecv = "server.Queries"
	}
	// LHS introduces a new variable (ownerVar) alongside err. Go's
	// short-declaration `:=` is required when at least one LHS name is
	// new, regardless of whether err was already declared in the
	// enclosing tx preamble. Using `=` would yield
	// `undefined: owner<Resource>` at compile time. (BUG-029)
	ownerAssign := g.assignOp(true)
	lines := []string{
		fmt.Sprintf("%s, err %s %s.%s(ctx, %s)", ownerVar, ownerAssign, queriesRecv, queryName, ridExpr),
		"if err != nil {",
		fmt.Sprintf("\t%s(\"handler: %s\", \"op\", %q, \"status\", %d, \"err\", err)", logLevelFuncForStatus(status), logTagForStatus(status), g.FuncName, status),
		fmt.Sprintf("\treturn api.%s%dJSONResponse{Error: %q, Code: strPtr(%q)}, nil", g.FuncName, status, msg, neutralCode(status)),
		"}",
	}
	ownersExpr := fmt.Sprintf(
		"map[string]map[string]any{%q: {fmt.Sprint(%s): %s}}",
		mapping.Resource, ridExpr, ownerVar)
	return lines, ownersExpr, []string{`"fmt"`}
}
