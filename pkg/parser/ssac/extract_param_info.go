//ff:func feature=ssac-parse type=parser control=iteration dimension=1
//ff:what AST 함수 선언에서 파라미터 정보를 추출
package ssac

import "go/ast"

// extractParamInfo는 AST 함수 선언에서 파라미터 정보를 추출한다.
func extractParamInfo(fn *ast.FuncDecl) *ParamInfo {
	if fn.Type.Params == nil {
		return nil
	}
	var param *ParamInfo
	for _, p := range fn.Type.Params.List {
		if len(p.Names) > 0 {
			param = &ParamInfo{
				TypeName: exprToString(p.Type),
				VarName:  p.Names[0].Name,
			}
		}
	}
	return param
}
