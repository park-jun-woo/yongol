//ff:func feature=validate type=rule control=iteration dimension=3 topic=ssac-structural
//ff:what S-50 — request.field가 OpenAPI request schema에 존재

package ssac

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s50OpenAPIRequest validates S-50: every request.<field> referenced in Args
// must be a declared field on the OpenAPI request schema for the operation.
//
// 관계: XOS-66 (openapi_ssac/) 는 대칭이 아닌 상위 제약 — 참조된 필드가
// requestBody.required 배열에도 포함되어야 한다. 본 S-50 은 단순 존재(schema
// properties) 만 검증하고 XOS-66 이 required-list 추가 검증을 담당한다.
// 같은 위반(필드 자체가 schema 에 없음)이면 S-50 이 먼저 ERROR 를 내고,
// XOS-66 은 required 누락만 별도 신호로 잡는다.
func s50OpenAPIRequest(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	g := fs.Ground()
	if g == nil {
		return nil
	}
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		if fn.Subscribe != nil {
			continue
		}
		fields, ok := g.Lookup["OpenAPI.request."+fn.Name]
		if !ok {
			continue
		}
		for _, seq := range fn.Sequences {
			for _, arg := range seq.Args {
				if arg.Source != "request" || arg.Field == "" {
					continue
				}
				if fields[arg.Field] {
					continue
				}
				diags = append(diags, diagnostic.Diagnostic{
					File:    fn.FileName,
					Line:    seq.Line,
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelError,
					Message: fmt.Sprintf("[S-50] request.%s not in OpenAPI request schema", arg.Field),
					Advice:  fmt.Sprintf("OpenAPI operationId %q 의 requestBody 에 필드 %q 를 추가하세요", fn.Name, arg.Field),
				})
			}
		}
	}
	return diags
}
