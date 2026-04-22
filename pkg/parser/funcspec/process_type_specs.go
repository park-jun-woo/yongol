//ff:func feature=funcspec type=parser control=iteration dimension=1
//ff:what GenDecl에서 Request/Response 구조체 필드를 추출한다
package funcspec

import (
	"go/ast"
)

func processTypeSpecs(d *ast.GenDecl, spec *FuncSpec, expectedRequest, expectedResponse string) {
	for _, s := range d.Specs {
		ts, ok := s.(*ast.TypeSpec)
		if !ok {
			continue
		}
		st, ok := ts.Type.(*ast.StructType)
		if !ok {
			continue
		}
		fields := extractFields(st)
		if ts.Name.Name == expectedRequest {
			spec.RequestFields = fields
		} else if ts.Name.Name == expectedResponse {
			spec.ResponseFields = fields
		}
	}
}
