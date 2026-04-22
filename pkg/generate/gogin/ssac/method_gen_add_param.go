//ff:func feature=gen-gogin type=util control=selection
//ff:what methodGen.addParam — ParameterRef 위치(path/query)에 맞게 methodGen 맵에 등록

package ssac

import "github.com/getkin/kin-openapi/openapi3"

// addParam registers a single parameter into PathParams or QueryParams
// with the metadata needed by mapValue (enum alias names, int64 format, etc.).
func (g *methodGen) addParam(p *openapi3.ParameterRef, operationId string) {
	if p == nil || p.Value == nil {
		return
	}
	switch p.Value.In {
	case "path":
		g.PathParams[p.Value.Name] = true
	case "query":
		g.QueryParams[p.Value.Name] = buildQueryParam(p, operationId)
	}
}
