//ff:func feature=gen-gogin type=util control=sequence
//ff:what ownershipLookup — emit OwnerLookup<Res> call + Owners literal for buildAuth (pgtypex bridge)

package ssac

import (
	"fmt"
	"strings"

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

	pkCol := g.lookupResourcePKColumn(mapping.Resource)
	alreadyPgtype := !strings.HasPrefix(rawRID, "request.")
	sqlcArg, imports := resolvePKSqlcArg(pkCol, ridExpr, alreadyPgtype)

	ownerAssign := g.assignOp(true)
	lines := []string{
		fmt.Sprintf("%s, err %s %s.%s(ctx, %s)", ownerVar, ownerAssign, queriesRecv, queryName, sqlcArg),
		"if err != nil {",
		fmt.Sprintf("\t%s(\"handler: %s\", \"op\", %q, \"status\", %d, \"err\", err)", logLevelFuncForStatus(status), logTagForStatus(status), g.FuncName, status),
		fmt.Sprintf("\treturn api.%s%dJSONResponse{Error: %q, Code: %q}, nil", g.FuncName, status, msg, neutralCode(status)),
		"}",
	}

	ownerKeyExpr := "fmt.Sprint(" + ridExpr + ")"
	if pkCol != nil && isUUIDColumn(pkCol) {
		if alreadyPgtype {
			ownerKeyExpr = "pgtypex.UUIDToString(" + ridExpr + ")"
		} else {
			ownerKeyExpr = "pgtypex.UUIDToString(pgtypex.ToPgUUID(" + ridExpr + "))"
		}
		imports = append(imports, `"github.com/park-jun-woo/ssac/pkg/pgtypex"`)
	} else {
		imports = append(imports, `"fmt"`)
	}
	ownersExpr := fmt.Sprintf(
		"map[string]map[string]any{%q: {%s: %s}}",
		mapping.Resource, ownerKeyExpr, ownerVar)
	return lines, ownersExpr, imports
}
