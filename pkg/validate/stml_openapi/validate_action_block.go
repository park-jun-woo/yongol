//ff:func feature=validate type=rule control=iteration dimension=1 topic=stml-openapi
//ff:what validateActionBlock — data-action 블록의 operationId·메서드·필드·컴포넌트 검증

package stml_openapi

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// validateActionBlock checks a data-action block against the OpenAPI spec.
// TM-02 (operationId), TM-03 (GET method), TM-04 (params), TM-05 (fields),
// TM-09 (component).
func validateActionBlock(a stml.ActionBlock, file string, opMap map[string]operationEntry, fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic

	entry, ok := opMap[a.OperationID]
	if !ok {
		// TM-02: data-action operationId not found
		diags = append(diags, tm02ActionOpNotFound(file, a.OperationID))
		return diags
	}

	// TM-03: data-action references a GET endpoint
	if entry.method == "GET" {
		diags = append(diags, tm03ActionGetMethod(file, a.OperationID))
	}

	// TM-04: params
	diags = append(diags, tm04Params(a.Params, a.OperationID, file, entry)...)

	// TM-05: fields / TM-19: object(map) field bound to a plain text input
	reqFields := requestBodyFields(entry.op)
	reqFieldTypes := requestBodyFieldTypes(entry.op)
	for _, f := range a.Fields {
		// Skip component references (handled by TM-09 below)
		if strings.HasPrefix(f.Tag, "data-component:") {
			comp := strings.TrimPrefix(f.Tag, "data-component:")
			diags = append(diags, tm09Component(comp, file, fs)...)
			continue
		}
		if _, ok := reqFields[f.Name]; !ok {
			diags = append(diags, tm05FieldNotFound(file, a.OperationID, f.Name))
			continue
		}
		// TM-19: object(map) type request field cannot be captured by a single
		// text input. There is no key-value widget yet, so any object-typed
		// data-field binding is flagged.
		if reqFieldTypes[f.Name] == "object" {
			diags = append(diags, tm19MapFieldTextInput(file, a.OperationID, f.Name))
		}
	}

	return diags
}
