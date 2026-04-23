//ff:func feature=validate type=rule control=iteration dimension=1 topic=ssac-openapi
//ff:what XOS-82 — OpenAPI 에 선언된 2xx 중 yongol 이 emit 하지 않는 것이 있음

package openapi_ssac

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	yopenapi "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xos82UnreachableSuccessStatus validates XOS-82: the OpenAPI operation
// declares multiple 2xx responses but yongol emits only the one selected
// by DeriveSuccessStatus. The remaining 2xx codes are unreachable from
// the generated handler. Authors may have pre-declared codes for
// forward-compat (e.g. future `@upsert` that returns either 200 or 201);
// WARNING keeps the operation visible without blocking codegen. The
// warning also nudges authors to trim genuinely dead declarations.
func xos82UnreachableSuccessStatus(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs.OpenAPIDoc == nil {
		return nil
	}
	opMap := buildOperationMethodMap(fs.OpenAPIDoc)
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		if !hasResponseSequence(fn) {
			continue
		}
		entry, ok := opMap[fn.Name]
		if !ok {
			continue
		}
		selected := yopenapi.DeriveSuccessStatus(entry.Op, entry.Method)
		if selected == 0 {
			continue
		}
		declared := yopenapi.Declared2xx(entry.Op)
		if len(declared) <= 1 {
			continue
		}
		unreachable := make(map[int]bool, len(declared))
		for code := range declared {
			if code != selected {
				unreachable[code] = true
			}
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:  fn.FileName,
			Line:  fn.Line,
			Phase: diagnostic.PhaseValidate,
			Level: diagnostic.LevelWarning,
			Message: fmt.Sprintf(
				"[XOS-82] operation %s declares 2xx %v but only %d is reachable from SSaC",
				fn.Name, sortedKeys(declared), selected),
			Advice: fmt.Sprintf("Either remove the unused 2xx declarations %v or extend SSaC to emit them",
				sortedKeys(unreachable)),
		})
	}
	return diags
}
