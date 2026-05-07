//ff:func feature=gen-gogin type=util control=sequence
//ff:what ownershipLookup — emit OwnerLookup<Res> call + Owners literal for buildAuth (pgtypex bridge)

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/gogin/types"
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
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

	// Resolve PK column to apply InsertExpr wrap (UUID → pgtypex.ToPgUUID)
	pkCol := g.lookupResourcePKColumn(mapping.Resource)
	sqlcArg := ridExpr
	var imports []string
	if pkCol != nil {
		binding := types.MapPGType(*pkCol)
		if binding.InsertExpr != "" && binding.InsertExpr != "{var}" {
			sqlcArg = types.Expand(binding.InsertExpr, "", "", ridExpr)
			for _, imp := range binding.Imports {
				if strings.Contains(imp, "pgtypex") {
					imports = append(imports, `"`+imp+`"`)
				}
			}
		}
	}

	ownerAssign := g.assignOp(true)
	lines := []string{
		fmt.Sprintf("%s, err %s %s.%s(ctx, %s)", ownerVar, ownerAssign, queriesRecv, queryName, sqlcArg),
		"if err != nil {",
		fmt.Sprintf("\t%s(\"handler: %s\", \"op\", %q, \"status\", %d, \"err\", err)", logLevelFuncForStatus(status), logTagForStatus(status), g.FuncName, status),
		fmt.Sprintf("\treturn api.%s%dJSONResponse{Error: %q, Code: strPtr(%q)}, nil", g.FuncName, status, msg, neutralCode(status)),
		"}",
	}

	// Owners map key: UUID → pgtypex.UUIDToString, others → fmt.Sprint
	ownerKeyExpr := "fmt.Sprint(" + ridExpr + ")"
	if pkCol != nil && isUUIDColumn(pkCol) {
		ownerKeyExpr = "pgtypex.UUIDToString(" + ridExpr + ")"
		imports = append(imports, `"github.com/park-jun-woo/ssac/pkg/pgtypex"`)
	} else {
		imports = append(imports, `"fmt"`)
	}
	ownersExpr := fmt.Sprintf(
		"map[string]map[string]any{%q: {%s: %s}}",
		mapping.Resource, ownerKeyExpr, ownerVar)
	return lines, ownersExpr, imports
}

// lookupResourcePKColumn resolves the PK (id) column for a resource name.
func (g *methodGen) lookupResourcePKColumn(resource string) *ddl.Column {
	return lookupDDLColumn(g.DDLTables, pascalCase(resource), "id")
}

// isUUIDColumn returns true when the column resolves to a UUID binding.
func isUUIDColumn(col *ddl.Column) bool {
	binding := types.MapPGType(*col)
	return binding.SqlcGoType == "pgtype.UUID"
}
