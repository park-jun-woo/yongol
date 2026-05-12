//ff:func feature=gen-react type=generator control=sequence
//ff:what writeApiClientEntry — api.<OperationID>(args) 엔트리 1건 방출 (path/query/body 분기)

package react

import (
	"fmt"
	"strings"
)

// writeApiClientEntry emits a single api.<OperationID>(args) entry.
// Path params are lifted out of `args` into the openapi-fetch `path` option;
// remaining properties flow as query (GET) or body (mutation).
//
// The wrapper signature uses Req<K>/Res<K> generics for type safety while
// casting through `as any` internally to bypass openapi-fetch path-type
// mismatches (controlled escape hatch).
func writeApiClientEntry(b *strings.Builder, ep endpoint) {
	method := strings.ToUpper(ep.method)
	pathLit := ep.path
	opQ := fmt.Sprintf("'%s'", ep.opID) // quoted operationId for type arg

	// --- function signature: typed Req<K> ---
	// GET without path params: all query keys are optional, so args itself is optional.
	optionalMark := ""
	if method == "GET" && len(ep.pathParams) == 0 {
		optionalMark = "?"
	}
	b.WriteString("  ")
	b.WriteString(ep.opID)
	b.WriteString(fmt.Sprintf(": (args%s: Req<%s>) => {\n", optionalMark, opQ))

	if len(ep.pathParams) > 0 {
		// Extract path params by name.
		b.WriteString("    const a = args as any\n")
		b.WriteString("    const path: Record<string, any> = {}\n")
		for _, pp := range ep.pathParams {
			b.WriteString(fmt.Sprintf("    if (a && a['%s'] !== undefined) path['%s'] = a['%s']\n", pp, pp, pp))
		}
	}
	if method == "GET" {
		if len(ep.pathParams) > 0 {
			b.WriteString("    const query: Record<string, any> = {}\n")
			b.WriteString("    if (a) {\n")
			b.WriteString("      for (const [k, v] of Object.entries(a)) {\n")
			b.WriteString("        if (v == null) continue\n")
			b.WriteString("        if (!(k in path)) query[k] = v\n")
			b.WriteString("      }\n")
			b.WriteString("    }\n")
			b.WriteString(fmt.Sprintf("    return client.GET('%s', { params: { path, query } } as any).then(r => r.data as Res<%s>)\n", pathLit, opQ))
		} else {
			b.WriteString(fmt.Sprintf("    return client.GET('%s', { params: { query: args ?? {} } } as any).then(r => r.data as Res<%s>)\n", pathLit, opQ))
		}
	} else {
		verbCall := "POST"
		switch method {
		case "PUT", "PATCH", "DELETE":
			verbCall = method
		}
		if len(ep.pathParams) > 0 {
			b.WriteString("    const body: Record<string, any> = {}\n")
			b.WriteString("    if (a) {\n")
			b.WriteString("      for (const [k, v] of Object.entries(a)) {\n")
			b.WriteString("        if (!(k in path)) body[k] = v\n")
			b.WriteString("      }\n")
			b.WriteString("    }\n")
			b.WriteString(fmt.Sprintf("    return client.%s('%s', { params: { path }, body } as any).then(r => r.data as Res<%s>)\n", verbCall, pathLit, opQ))
		} else {
			b.WriteString(fmt.Sprintf("    return client.%s('%s', { body: args ?? {} } as any).then(r => r.data as Res<%s>)\n", verbCall, pathLit, opQ))
		}
	}
	b.WriteString("  },\n")
}
