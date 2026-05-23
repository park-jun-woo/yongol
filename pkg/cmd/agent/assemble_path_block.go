//ff:func feature=agent type=helper control=iteration dimension=1
//ff:what assemblePathBlock — parameters + requestBody + responses를 OpenAPI path 블록으로 조립

package agent

import "github.com/park-jun-woo/yongol/pkg/parser/features"

func assemblePathBlock(feat features.Feature, params []any, reqBody map[string]any, schema200 map[string]any, errorResps map[string]any) map[string]any {
	method := httpMethodFromOp(feat.Op)

	op := map[string]any{
		"operationId": feat.Op,
		"summary":     feat.Desc,
	}

	if !feat.Public {
		op["security"] = []map[string][]string{
			{"bearerAuth": {}},
		}
	}

	if len(params) > 0 {
		op["parameters"] = params
	}

	if reqBody != nil && needsRequestBody(method) {
		op["requestBody"] = reqBody
	}

	responses := make(map[string]any)
	responses["200"] = map[string]any{
		"description": "OK",
		"content": map[string]any{
			"application/json": map[string]any{
				"schema": schema200,
			},
		},
	}
	for code, resp := range errorResps {
		responses[code] = resp
	}
	op["responses"] = responses

	return map[string]any{method: op}
}
