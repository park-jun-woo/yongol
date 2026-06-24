//ff:func feature=validate type=rule control=iteration dimension=1 topic=ssac-openapi
//ff:what XOS-80 — HTTP method 관례 성공 상태가 OpenAPI responses 에 없음

package openapi_ssac

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
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
	opMap := buildOperationMethodMapAll(fs)
	if len(opMap) == 0 {
		return nil
	}
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		diags = append(diags, xos80CheckFunc(fn, opMap)...)
	}
	return diags
}
