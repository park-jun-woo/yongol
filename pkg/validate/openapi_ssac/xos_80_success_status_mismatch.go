//ff:func feature=validate type=rule control=iteration dimension=1 topic=ssac-openapi
//ff:what XOS-80 — HTTP method 관례 성공 상태가 OpenAPI responses 에 없음

package openapi_ssac

import (
	"fmt"
	"sort"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	yopenapi "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xos80SuccessStatusMismatch validates XOS-80: for each SSaC service func
// with an @response, verify that the HTTP-method-conventional 2xx status
// (POST→201, PUT/PATCH→200, DELETE→204, GET→200) is actually declared on
// the matching OpenAPI operation. yongol generate picks this status at
// codegen time; a mismatch produces an `undefined: api.<Op><Code>JSONResponse`
// build failure that was previously only caught at `go build`.
//
// We accept any declared 2xx — DeriveSuccessStatus also falls back to the
// lowest declared 2xx — but if neither the conventional code nor any
// other 2xx sits in responses, XOS-22 already fires (no 2xx). XOS-80
// specifically targets the "declared non-conventional 2xx that yongol
// does not know how to emit" case, which for now is unreachable given
// DeriveSuccessStatus's fallback; keep the diagnostic wired in so future
// stricter policy (e.g. refuse fallback selection) activates without a
// new rule ID.
func xos80SuccessStatusMismatch(fs *yongol.Fullstack) []diagnostic.Diagnostic {
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
		if yopenapi.DeriveSuccessStatus(entry.Op, entry.Method) != 0 {
			continue
		}
		// Unreachable as long as XOS-22 is also installed (it fires on
		// the empty-2xx case first). Kept so XOS-80 has a well-defined
		// meaning if XOS-22 is ever relaxed.
		declared := sortedKeys(yopenapi.Declared2xx(entry.Op))
		diags = append(diags, diagnostic.Diagnostic{
			File:  fn.FileName,
			Line:  fn.Line,
			Phase: diagnostic.PhaseValidate,
			Level: diagnostic.LevelError,
			Message: fmt.Sprintf(
				"[XOS-80] operation %s (%s) cannot derive a success status — declared 2xx: %v",
				fn.Name, entry.Method, declared),
			Advice: "Declare a conventional 2xx response for this method (POST→201, PUT/PATCH→200, DELETE→204, GET→200)",
		})
	}
	return diags
}

// sortedKeys returns the int keys of m in ascending order — used only by
// diagnostic messages so the output is deterministic across runs.
func sortedKeys(m map[int]bool) []int {
	out := make([]int, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	sort.Ints(out)
	return out
}
