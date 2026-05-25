//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what methodGen.buildFieldResponse — 필드 맵을 typed response (db→api 변환 포함, pgtype 자동 변환) 로 렌더링

package ssac

import (
	"fmt"
	"sort"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

// buildFieldResponse generates typed response with db→api conversion.
// SSaC fields: { workflow: updated, count: cc.Count }
// OpenAPI RespFields tells us the api type for each field.
//
// convert<Model> now returns (*api.Model, error) so $ref-typed fields
// have their conversion hoisted to local variables before the struct
// literal and any error is propagated as nil,err from the handler
// (BUG-003 / BUG-005 response direction).
//
// PhaseG02 — when a dotted expression (e.g. user.Name) accesses a column
// whose DDL type maps to a pgtype wrapper (pgtype.Text, pgtype.Int8,
// pgtype.Timestamptz, etc.), the mapped expression is replaced with the
// pgtypex bridge call so the generated assignment compiles against the
// oapi-codegen response struct. Required imports are collected and
// returned alongside the lines.
func (g *methodGen) buildFieldResponse(fields map[string]string) ([]string, []string) {
	var lines []string
	var imports []string

	// 필드 값을 Go 표현식으로 변환 (request.id → request.Id 등)
	mapped := make(map[string]string, len(fields))
	for k, v := range fields {
		mapped[k] = g.mapValue(v)
	}

	// Phase002-ManifestRef: resolve manifest.* references to Go literals.
	// "manifest.auth.accessTokenTTL" → "900" (int64 seconds from "15m").
	// mapValue does not touch manifest.* because its prefix is not "request",
	// so the original SSaC value is preserved in mapped[k]. We resolve here
	// and replace with the Go literal so downstream rendering treats it as
	// a numeric literal (isLiteral → true, isIntegerLiteralStr → true).
	for k, v := range fields {
		if !strings.HasPrefix(v, "manifest.") {
			continue
		}
		refPath := strings.TrimPrefix(v, "manifest.")
		if rv, ok := manifest.ResolveRef(g.Manifest, refPath); ok {
			mapped[k] = rv.GoLit
		}
	}

	keys := make([]string, 0, len(fields))
	for k := range fields {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// PhaseG02: resolve pgtype conversions for dotted field accesses.
	// When varExpr is e.g. "user.Name" and the underlying DDL column is
	// pgtype.Text (nullable TEXT), replace the expression with
	// pgtypex.FromPgTextPtr(user.Name) and collect imports. The converted
	// expression is stored in pgtypeConverted so renderResponseFieldHoisted
	// can use it for correct struct literal assignment.
	pgtypeConverted := make(map[string]string)
	for _, jsonName := range keys {
		varExpr := mapped[jsonName]
		if convExpr, convImports := g.resolvePgtypeFieldExpr(varExpr); convExpr != "" {
			pgtypeConverted[jsonName] = convExpr
			imports = append(imports, convImports...)
		}
	}

	// Pre-pass: hoist any $ref convert<Type>/convert<Type>List call into
	// a local variable so the per-call error is reachable. scalarLocal
	// keyed by jsonName → local var name; same for listLocal.
	scalarLocal := make(map[string]string)
	listLocal := make(map[string]string)
	for _, jsonName := range keys {
		rf, ok := g.RespFields[jsonName]
		if !ok || rf.RefType == "" {
			continue
		}
		varExpr := mapped[jsonName]
		local := lowerFirst(pascalCase(jsonName)) + "Converted"
		if rf.IsArray {
			listLocal[jsonName] = local
			lines = append(lines,
				fmt.Sprintf("%s, err := convert%sList(%s)", local, rf.RefType, varExpr),
				"if err != nil { return nil, err }",
			)
		} else {
			scalarLocal[jsonName] = local
			lines = append(lines,
				fmt.Sprintf("%s, err := convert%s(%s)", local, rf.RefType, varExpr),
				"if err != nil { return nil, err }",
			)
		}
	}

	// g.SuccessStatus is HTTP-method-derived at extract time (BUG-004) —
	// POST → 201, DELETE → 204, etc. instead of the prior hardcoded 200.
	lines = append(lines, fmt.Sprintf("return api.%s%dJSONResponse{",
		g.FuncName, g.SuccessStatus))

	for _, jsonName := range keys {
		varExpr := mapped[jsonName]
		if conv, ok := pgtypeConverted[jsonName]; ok {
			lines = append(lines, g.renderPgtypeField(jsonName, varExpr, conv))
			continue
		}
		lines = append(lines, g.renderResponseFieldHoisted(jsonName, varExpr, scalarLocal, listLocal))
	}
	lines = append(lines, "}, nil")
	return lines, imports
}
