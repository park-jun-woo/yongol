//ff:func feature=validate type=rule control=iteration dimension=3 topic=ssac-structural
//ff:what S-59 — 변수.필드 접근에서 필드가 변수 타입의 실제 필드인지 검증

package ssac

import (
	"fmt"

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
					diags = append(diags, diagnostic.Diagnostic{
						File:    fn.FileName,
						Line:    seq.Line,
						Phase:   diagnostic.PhaseValidate,
						Level:   diagnostic.LevelError,
						Message: fmt.Sprintf("[S-59] %s.%s: field %q does not exist on variable type", arg.Source, arg.Field, arg.Field),
						Advice:  fmt.Sprintf("변수 %s 의 타입 필드 중 정확한 이름으로 수정하세요", arg.Source),
					})
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
					diags = append(diags, diagnostic.Diagnostic{
						File:    fn.FileName,
						Line:    seq.Line,
						Phase:   diagnostic.PhaseValidate,
						Level:   diagnostic.LevelError,
						Message: fmt.Sprintf("[S-59] %s.%s: field %q does not exist on variable type", src, fld, fld),
						Advice:  fmt.Sprintf("변수 %s 의 타입 필드 중 정확한 이름으로 수정하세요", src),
					})
				}
			}
		}
	}
	return diags
}
