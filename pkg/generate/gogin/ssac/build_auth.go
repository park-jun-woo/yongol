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

	// `assign` selects := vs = for the final authz.Check call. When the
	// ownership branch below runs it re-evaluates assignOp after emitting
	// the owner lookup (which itself declares err via :=), so `assign`
	// must remain reactive to methodGen's FirstErr state. (BUG-029)
	assign := g.assignOp(false)

	mapping := findOwnershipMapping(g.Ownerships, seq.Resource)
	var lines []string
	ownersExpr := "nil"
	imports := []string{`"github.com/park-jun-woo/ssac/pkg/authz"`, `"log/slog"`}

	// Phase005 (BUG-033): for creation-form endpoints the ResourceID is
	// statically zero — the resource does not exist yet. Calling
	// OwnerLookup<Res>(ctx, 0) always yields sql.ErrNoRows → 403. Skip
	// the owner-lookup injection (and leave Owners nil) so the Rego
	// policy evaluates role-only rules that gate the create action.
	// Update/Delete/Get still carry a non-zero ResourceID and keep the
	// lookup. Detection is static — based on the SSaC AST expression
	// (absent, "0", empty string, nil, null).
	rawRID, hasRID := seq.Inputs["ResourceID"]
	if mapping != nil && hasRID && !isResourceIDZero(rawRID) {
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
		// short-declaration `:=` is required — and permitted — when at
		// least one LHS name is new, regardless of whether err was
		// already declared in the enclosing tx preamble. Using `=`
		// would yield `undefined: owner<Resource>` at compile time.
		// (BUG-029)
		ownerAssign := g.assignOp(true)
		lines = append(lines,
			fmt.Sprintf("%s, err %s %s.%s(ctx, %s)", ownerVar, ownerAssign, queriesRecv, queryName, ridExpr),
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

// isResourceIDZero reports whether the ResourceID expression pulled
// from a `@auth` sequence's Inputs map is statically zero — i.e. the
// handler is a creation form and no resource exists yet. Detection is
// by inspection of the SSaC AST expression; runtime zeros (e.g. a
// variable that happens to resolve to 0) are out of scope. Matched
// literals: empty string, `0`, `""`, `nil`, `null` (case-insensitive),
// with surrounding whitespace ignored.
//
// The caller (buildAuth) combines this with a presence check: a
// missing ResourceID key also counts as zero. See Phase005 (BUG-033).
func isResourceIDZero(expr string) bool {
	s := strings.TrimSpace(expr)
	if s == "" {
		return true
	}
	switch strings.ToLower(s) {
	case "0", `""`, "''", "nil", "null":
		return true
	}
	return false
}
