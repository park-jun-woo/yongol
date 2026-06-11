//ff:func feature=validate type=rule control=iteration dimension=3 topic=ssac-structural
//ff:what s59DottedField — S-59 검증: variable.field 참조 시 필드가 변수 타입에 실재해야 함

package ssac

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s59DottedField validates S-59: when an SSaC sequence references a variable
// member (e.g. `wf.Status`), the member must be a field of the variable's
// type. Relies on Ground.Schemas["SSaC.var.<func>.<var>"] populated by
// populateSSaCSymbols.
//
// Skip conditions:
//   - source is a reserved keyword (request, query, currentUser, message)
//   - variable schema not registered (non-DDL type; e.g. @call Response)
//   - field is empty
func s59DottedField(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	g := fs.Ground()
	if g == nil {
		return nil
	}
	reserved := map[string]bool{
		"request": true, "query": true, "currentUser": true, "message": true,
	}
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			for _, arg := range seq.Args {
				if arg.Source == "" || arg.Field == "" || reserved[arg.Source] {
					continue
				}
				schema, ok := g.Schemas["SSaC.var."+fn.Name+"."+arg.Source]
				if !ok {
					continue
				}
				if !containsField(schema, arg.Field) {
					diags = append(diags, s59FieldDiag(fn.FileName, seq.Line, arg.Source, arg.Field, schema))
				}
			}
			// seq.Inputs values: "source.field" — same check
			for _, val := range seq.Inputs {
				src, fld := parseDottedValue(val)
				if src == "" || fld == "" || reserved[src] {
					continue
				}
				schema, ok := g.Schemas["SSaC.var."+fn.Name+"."+src]
				if !ok {
					continue
				}
				if !containsField(schema, fld) {
					diags = append(diags, s59FieldDiag(fn.FileName, seq.Line, src, fld, schema))
				}
			}
		}
	}
	return diags
}
