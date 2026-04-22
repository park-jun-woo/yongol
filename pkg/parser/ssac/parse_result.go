//ff:func feature=ssac-parse type=parser control=sequence topic=response
//ff:what "Type var" 또는 "[]Type var" 결과 바인딩 파싱
package ssac

import "strings"

// parseResult는 "Type var" 또는 "[]Type var"를 파싱한다.
func parseResult(lhs string) *Result {
	lhs = strings.TrimSpace(lhs)
	parts := strings.Fields(lhs)
	if len(parts) != 2 {
		return nil
	}
	typeName := parts[0]
	r := &Result{Var: parts[1]}

	// Page[Gig] → Wrapper="Page", Type="Gig"
	// Cursor[Gig] → Wrapper="Cursor", Type="Gig"
	if bracketIdx := strings.IndexByte(typeName, '['); bracketIdx > 0 {
		if strings.HasSuffix(typeName, "]") {
			r.Wrapper = typeName[:bracketIdx]
			r.Type = typeName[bracketIdx+1 : len(typeName)-1]
			return r
		}
	}

	r.Type = typeName
	return r
}
