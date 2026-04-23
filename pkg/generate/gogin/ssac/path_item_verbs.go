//ff:func feature=gen-gogin type=util control=sequence
//ff:what pathItemVerbs — PathItem에서 (verb,Operation) 페어 배열을 생성
package ssac

import "github.com/getkin/kin-openapi/openapi3"

func pathItemVerbs(pathItem *openapi3.PathItem) []verbOp {
	return []verbOp{
		{"GET", pathItem.Get},
		{"POST", pathItem.Post},
		{"PUT", pathItem.Put},
		{"DELETE", pathItem.Delete},
		{"PATCH", pathItem.Patch},
	}
}
