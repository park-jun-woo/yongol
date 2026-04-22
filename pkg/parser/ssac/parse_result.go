//ff:func feature=ssac-parse type=parser control=sequence topic=response
//ff:what parseResult — parses a "Type var" or "[]Type var" result binding
package ssac

import "strings"

// parseResult parses a result binding of the form "Type var" or "[]Type var".
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
