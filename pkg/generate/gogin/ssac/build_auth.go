//ff:func feature=gen-gogin type=util control=sequence
//ff:what buildAuth — @auth 시퀀스 빌더 (OwnerLookup<Resource> + Owners 맵 + authz.Check)

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/rego"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// buildAuth emits the authz.Check call for an `@auth` sequence. The entire
// *model.UserClaim struct is forwarded as the `Claim` field; rego.Input
// json-marshals it using the JSON tags produced by generate_user_claim.go
// into the `input.claims` object expected by the policy.
//
// Phase003 (ssac/purify) — ownership lookups.
//
// When the project's Rego declares `@ownership <resource>: <table>.<col>`
// mappings, this emitter:
//
//  1. Locates the mapping whose Resource matches seq.Resource.
//  2. Resolves the resource-id expression from seq.Inputs["ResourceID"] via
//     g.mapValue (falls back to `0` as an int64 literal if absent — the
//     authz policy then sees no matching owner and denies, preserving
//     safety).
//  3. Emits `<owner>, err := qtx.OwnerLookup<Pascal(resource)>(ctx, <rid>)`
//     when the handler has an active qtx (mutating seq present); otherwise
//     falls back to `server.Queries.<...>`.
//  4. Builds the `authz.CheckRequest.Owners` literal —
//       map[string]map[string]any{
//         "<resource>": {fmt.Sprint(<rid>): <owner>},
//       }
//     so the owner-id keeps its Go-native type (int64 / uuid / string)
//     and serialises into rego input as a matching JSON type. The
//     resource-id map key is stringified because OPA in-memory stores and
//     JSON objects require string keys; stringifying the *key* is the
//     canonical workaround recommended in plans/ssac/purify README's
//     "설계 결정" section.
//
// When no ownership mapping matches seq.Resource, the emitter skips the
// lookup altogether and emits Owners: nil — ssac/pkg/authz normalises that
// to an empty map, so policies that do not reference data.owners still
// evaluate correctly.
//
// Phase005 pgx/v5 refit removed the former authz.CheckRequest.Tx field.
// Handlers no longer thread a transaction through Check; the lookup runs
// on the enclosing qtx (MVCC snapshot aligned with in-tx writes) which
// resolves the Phase005 note about committed-only reads.
func (g *methodGen) buildAuth(seq ssacparser.Sequence) ([]string, []string) {
	status := resolveErrStatus("auth", seq.ErrStatus)
	msg := seq.Message
	if msg == "" {
		msg = neutralMessage(status)
	}
	extraFields := g.mapFields(seq.Inputs)
	assign := g.assignOp(false)

	mapping := findOwnershipMapping(g.Ownerships, seq.Resource)
	var lines []string
	ownersExpr := "nil"
	imports := []string{`"github.com/park-jun-woo/ssac/pkg/authz"`, `"log/slog"`}

	if mapping != nil {
		ridExpr := g.mapValue(seq.Inputs["ResourceID"])
		if ridExpr == "" {
			ridExpr = "0"
		}
		queryName := ownerLookupQueryName(mapping.Resource)
		ownerVar := "owner" + pascalCase(mapping.Resource)
		queriesRecv := "qtx"
		if !g.UseTx {
			queriesRecv = "server.Queries"
		}
		lines = append(lines,
			fmt.Sprintf("%s, err %s %s.%s(ctx, %s)", ownerVar, assign, queriesRecv, queryName, ridExpr),
			"if err != nil {",
			fmt.Sprintf("\t%s(\"handler: %s\", \"op\", %q, \"status\", %d, \"err\", err)", logLevelFuncForStatus(status), logTagForStatus(status), g.FuncName, status),
			fmt.Sprintf("\treturn api.%s%dJSONResponse{Error: %q, Code: strPtr(%q)}, nil", g.FuncName, status, msg, neutralCode(status)),
			"}",
		)
		ownersExpr = fmt.Sprintf(
			"map[string]map[string]any{%q: {fmt.Sprint(%s): %s}}",
			mapping.Resource, ridExpr, ownerVar)
		imports = append(imports, `"fmt"`)
		// After emitting the initial lookup, subsequent err assignments in
		// this block must reuse `=` (the first := already declared err).
		assign = g.assignOp(false)
	}

	// Compose the authz.Check literal. Owners is placed last so the
	// assembled string is stable regardless of mapping presence.
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
	lines = append(lines,
		fmt.Sprintf("_, err %s authz.Check(authz.CheckRequest{%s})", assign, strings.Join(checkFields, ", ")),
		"if err != nil {",
		fmt.Sprintf("\t%s(\"handler: %s\", \"op\", %q, \"status\", %d, \"err\", err)", logLevelFuncForStatus(status), logTagForStatus(status), g.FuncName, status),
		fmt.Sprintf("\treturn api.%s%dJSONResponse{Error: %q, Code: strPtr(%q)}, nil", g.FuncName, status, msg, neutralCode(status)),
		"}",
	)
	return lines, imports
}

// findOwnershipMapping returns the first parsed `@ownership` mapping whose
// Resource matches the @auth sequence's Resource. Nil when no mapping
// applies — the caller then omits the lookup.
func findOwnershipMapping(ownerships []rego.OwnershipMapping, resource string) *rego.OwnershipMapping {
	for i := range ownerships {
		if ownerships[i].Resource == resource {
			return &ownerships[i]
		}
	}
	return nil
}

// ownerLookupQueryName builds the canonical sqlc query name
// (`OwnerLookup<PascalResource>`) shared with XQP-30 validate. Keep in
// lockstep with pkg/validate/query_rego/xqp_30_owner_lookup_query.go.
func ownerLookupQueryName(resource string) string {
	return "OwnerLookup" + pascalCase(resource)
}
