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
func writeApiClientEntry(b *strings.Builder, ep endpoint) {
	method := strings.ToUpper(ep.method)
	pathLit := ep.path

	// args parameter signature — always optional so AI can call api.X().
	b.WriteString("  ")
	b.WriteString(ep.opID)
	b.WriteString(": (args?: Record<string, any>) => {\n")

	if len(ep.pathParams) > 0 {
		// Extract path params by name.
		b.WriteString("    const path: Record<string, any> = {}\n")
		for _, pp := range ep.pathParams {
			b.WriteString(fmt.Sprintf("    if (args && args['%s'] !== undefined) path['%s'] = args['%s']\n", pp, pp, pp))
		}
	}
	if method == "GET" {
		if len(ep.pathParams) > 0 {
			b.WriteString("    const query: Record<string, any> = {}\n")
			b.WriteString("    if (args) {\n")
			b.WriteString("      for (const [k, v] of Object.entries(args)) {\n")
			b.WriteString("        if (v == null) continue\n")
			b.WriteString("        if (!(k in path)) query[k] = v\n")
			b.WriteString("      }\n")
			b.WriteString("    }\n")
			b.WriteString(fmt.Sprintf("    return client.GET('%s', { params: { path, query } }).then(r => r.data)\n", pathLit))
		} else {
			b.WriteString(fmt.Sprintf("    return client.GET('%s', { params: { query: args ?? {} } }).then(r => r.data)\n", pathLit))
		}
	} else {
		verbCall := "POST"
		switch method {
		case "PUT", "PATCH", "DELETE":
			verbCall = method
		}
		if len(ep.pathParams) > 0 {
			b.WriteString("    const body: Record<string, any> = {}\n")
			b.WriteString("    if (args) {\n")
			b.WriteString("      for (const [k, v] of Object.entries(args)) {\n")
			b.WriteString("        if (!(k in path)) body[k] = v\n")
			b.WriteString("      }\n")
			b.WriteString("    }\n")
			b.WriteString(fmt.Sprintf("    return client.%s('%s', { params: { path }, body }).then(r => r.data)\n", verbCall, pathLit))
		} else {
			b.WriteString(fmt.Sprintf("    return client.%s('%s', { body: args ?? {} }).then(r => r.data)\n", verbCall, pathLit))
		}
	}
	b.WriteString("  },\n")
}
