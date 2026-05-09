//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-statemachine
//ff:what XSM-71 — @state input value type must be string-compatible (statemachine params are string)

package ssac_statemachine

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xsm71StateInputType validates XSM-71: every @state input value must
// resolve to a string-compatible Go type. Statemachine functions accept
// status parameters as string; passing pgtype.UUID or int64 causes a build
// failure that validate should catch early.
func xsm71StateInputType(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	g := fs.Ground()
	if g == nil {
		return nil
	}

	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if seq.Type != "state" {
				continue
			}
			for inputKey, inputValue := range seq.Inputs {
				sourceType := resolveStateInputType(g, fn.Name, inputValue)
				if sourceType == "" {
					continue
				}
				if !stateTypesCompatible(sourceType, "string") {
					diags = append(diags, diagnostic.Diagnostic{
						File:  fn.FileName,
						Line:  seq.Line,
						Phase: diagnostic.PhaseValidate,
						Level: diagnostic.LevelError,
						Message: "[XSM-71] @state input " + inputKey +
							" value type " + sourceType + " is not string-compatible" +
							" (statemachine parameter is string)",
						Advice: "Use a string-typed field (e.g. status TEXT column) " +
							"instead of a UUID or numeric column",
					})
				}
			}
		}
	}
	return diags
}

// resolveStateInputType resolves the Go type of a @state input value
// expression using Ground.Types. Duplicates the var.Field logic from
// ssac_func.resolveInputType to avoid cross-package dependency.
func resolveStateInputType(g *rule.Ground, funcName, value string) string {
	value = strings.TrimSpace(value)
	if value == "" {
		return ""
	}
	// literal check: quoted string → "string", numeric/bool/nil
	if t := inferStateLiteralType(value); t != "" {
		return t
	}
	// currentUser.Field
	if strings.HasPrefix(value, "currentUser.") {
		field := value[len("currentUser."):]
		return g.Types["Manifest.claim."+field]
	}
	// var.Field → SSaC.var → DDL.field
	if dot := strings.IndexByte(value, '.'); dot > 0 {
		varName := value[:dot]
		fieldName := value[dot+1:]
		modelType := g.Types["SSaC.var."+funcName+"."+varName]
		if modelType != "" {
			modelType = strings.TrimPrefix(strings.TrimPrefix(modelType, "[]"), "*")
			return g.Types["DDL.field."+modelType+"."+fieldName]
		}
	}
	if strings.ContainsAny(value, ".\"'") {
		return ""
	}
	return g.Types["SSaC.var."+funcName+"."+value]
}

// inferStateLiteralType returns the Go type of a literal value.
func inferStateLiteralType(value string) string {
	if len(value) >= 2 && value[0] == '"' && value[len(value)-1] == '"' {
		return "string"
	}
	if value == "true" || value == "false" {
		return "bool"
	}
	if value == "nil" {
		return "nil"
	}
	if len(value) > 0 && (value[0] >= '0' && value[0] <= '9' || value[0] == '-') {
		if strings.Contains(value, ".") {
			return "float64"
		}
		return "int"
	}
	return ""
}

// stateTypesCompatible reports whether actual can be assigned to expected.
// Mirrors ssac_func.TypesCompatible logic.
func stateTypesCompatible(actual, expected string) bool {
	a := strings.TrimPrefix(actual, "*")
	e := strings.TrimPrefix(expected, "*")
	if a == e {
		return true
	}
	if a == "nil" || e == "nil" {
		return true
	}
	return false
}
