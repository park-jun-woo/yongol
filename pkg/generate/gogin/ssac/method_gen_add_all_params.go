//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what methodGen.addAllParams — path-level + operation-level 파라미터를 모두 methodGen 에 등록

package ssac

import "github.com/getkin/kin-openapi/openapi3"

func (g *methodGen) addAllParams(pathItem *openapi3.PathItem, v verbOp, operationId string) {
	for _, p := range pathItem.Parameters {
		g.addParam(p, operationId)
	}
	for _, p := range v.op.Parameters {
		g.addParam(p, operationId)
	}
}
