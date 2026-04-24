//ff:func feature=validate type=rule control=iteration dimension=1 topic=states
//ff:what XSM-27 — stateful POST/PUT/DELETE 의 @state / @state-neutral 의도 선언 강제

package ssac_statemachine

import (
	"strings"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xsm27StateIntentDeclaration validates XSM-27: any POST/PUT/DELETE that
// targets a stateful resource (via a `{id}` path parameter) and reads that
// resource with `@get <Model>.FindByID({ID: request.id})` must either declare
// a `@state` guard or explicitly opt out with `// @state-neutral`.
//
// Fires WARNING when all the following hold:
//  1. The OpenAPI operation's path contains an `{id}`-ish parameter
//  2. The HTTP method is POST, PUT, or DELETE
//  3. The first path segment maps to a stateful resource (state diagram
//     exists and DDL DEFAULT matches the diagram's initial state — XDM-28 linkage)
//  4. The SSaC function reads that resource via `@get <StatefulModel>.FindByID({ID: request.id})`
//  5. Neither a `@state` sequence nor a `// @state-neutral` annotation is present
func xsm27StateIntentDeclaration(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil || fs.OpenAPIDoc == nil || fs.OpenAPIDoc.Paths == nil {
		return nil
	}
	g := fs.Ground()
	if g == nil {
		return nil
	}
	funcByName := buildFuncByName(fs.ServiceFuncs)
	var diags []diagnostic.Diagnostic
	for pathStr, item := range fs.OpenAPIDoc.Paths.Map() {
		if item == nil {
			continue
		}
		if !pathHasIDParam(pathStr) {
			continue
		}
		target := isStatefulResource(pathStr, fs.StateDiagrams, g)
		if target == nil {
			continue
		}
		for _, v := range []struct {
			method string
			op     *openapi3.Operation
		}{
			{"POST", item.Post},
			{"PUT", item.Put},
			{"DELETE", item.Delete},
		} {
			if v.op == nil || v.op.OperationID == "" {
				continue
			}
			fn, ok := funcByName[v.op.OperationID]
			if !ok {
				continue
			}
			if fn.StateNeutral {
				continue
			}
			if hasStateSequence(fn.Sequences) {
				continue
			}
			resultVar, ok := findByIDResultVar(fn.Sequences, target.Model)
			if !ok {
				continue
			}
			diags = append(diags, buildXsm27Diag(fn, target, resultVar))
		}
	}
	return diags
}

// pathHasIDParam reports whether an OpenAPI path template contains an {id}
// parameter. Lenient on case so `{ID}` / `{Id}` pass too.
func pathHasIDParam(path string) bool {
	for _, part := range strings.Split(path, "/") {
		if !strings.HasPrefix(part, "{") || !strings.HasSuffix(part, "}") {
			continue
		}
		name := strings.TrimSuffix(strings.TrimPrefix(part, "{"), "}")
		if strings.EqualFold(name, "id") {
			return true
		}
	}
	return false
}

// findByIDResultVar returns the bound variable name of the `@get
// <Model>.FindByID(...)` sequence and reports whether such a sequence was
// found. When the `@get` has no result variable the empty string is returned
// alongside ok=true, which still signals "resource is read" for detection
// purposes; callers pick a fallback name for diagnostics.
func findByIDResultVar(seqs []ssac.Sequence, model string) (string, bool) {
	if model == "" {
		return "", false
	}
	wanted := model + ".FindByID"
	for _, seq := range seqs {
		if seq.Type != "get" {
			continue
		}
		if seq.Model != wanted {
			continue
		}
		if seq.Result != nil {
			return seq.Result.Var, true
		}
		return "", true
	}
	return "", false
}

// buildXsm27Diag assembles the full WARNING with the A/B options and the
// self-loop transition hint derived from the diagram's initial state. The
// `resultVar` is the variable name from the @get FindByID sequence; when
// empty the lowercase resource name is used as a readable fallback.
func buildXsm27Diag(fn ssac.ServiceFunc, target *statefulTarget, resultVar string) diagnostic.Diagnostic {
	file := fn.FileName
	if file == "" {
		file = "ssac/" + fn.Name + ".ssac"
	}
	diagramID := ""
	initial := ""
	if target.Diagram != nil {
		diagramID = target.Diagram.ID
		initial = target.Diagram.InitialState
	}
	msg := "[XSM-27] " + fn.Name + ": state-dependent operation on stateful resource '" +
		target.Resource + "' is missing @state declaration"

	varName := resultVar
	if varName == "" {
		varName = strings.ToLower(target.Resource)
	}

	var b strings.Builder
	b.WriteString("Option A (state-dependent): add above `func ")
	b.WriteString(fn.Name)
	b.WriteString("() {}`\n")
	b.WriteString("    // @state ")
	b.WriteString(diagramID)
	b.WriteString(" {")
	b.WriteString(pascalCaseFromLower(target.StateColumn))
	b.WriteString(": ")
	b.WriteString(varName)
	b.WriteString(".")
	b.WriteString(pascalCaseFromLower(target.StateColumn))
	b.WriteString("} \"")
	b.WriteString(fn.Name)
	b.WriteString("\" \"Cannot ")
	b.WriteString(fn.Name)
	b.WriteString("\" 409\n")
	if diagramID != "" && initial != "" {
		b.WriteString("  If ")
		b.WriteString(fn.Name)
		b.WriteString(" is not declared as a transition in states/")
		b.WriteString(diagramID)
		b.WriteString(".md, add it (self-loop if no state change):\n")
		b.WriteString("    ")
		b.WriteString(initial)
		b.WriteString(" --> ")
		b.WriteString(initial)
		b.WriteString(": ")
		b.WriteString(fn.Name)
		b.WriteString("\n")
	}
	b.WriteString("Option B (state-neutral): if this operation truly does not depend on the resource's state, add above the function:\n")
	b.WriteString("    // @state-neutral")

	return diagnostic.Diagnostic{
		File:    file,
		Line:    fn.Line,
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelWarning,
		Message: msg,
		Advice:  b.String(),
	}
}
