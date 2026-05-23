//ff:func feature=agent type=helper control=sequence
//ff:what buildErrorResponses — feature 속성 기반 에러 응답 맵 생성

package agent

import "github.com/park-jun-woo/yongol/pkg/parser/features"

func buildErrorResponses(feat features.Feature) map[string]any {
	errRef := map[string]any{
		"description": "Error",
		"content": map[string]any{
			"application/json": map[string]any{
				"schema": map[string]any{
					"$ref": "#/components/schemas/Error",
				},
			},
		},
	}

	method := httpMethodFromOp(feat.Op)
	resps := make(map[string]any)

	if feat.Public {
		resps["400"] = copyMap(errRef)
		resps["500"] = copyMap(errRef)
	} else if method == "delete" {
		resps["401"] = copyMap(errRef)
		resps["403"] = copyMap(errRef)
		resps["404"] = copyMap(errRef)
	} else {
		resps["401"] = copyMap(errRef)
		resps["403"] = copyMap(errRef)
		resps["404"] = copyMap(errRef)
		resps["500"] = copyMap(errRef)
	}

	return resps
}
