//ff:func feature=gen-react type=generator control=sequence
//ff:what writeApiClientEntry — api.<OperationID>(args) 엔트리 1건 방출 (path/query/body 분기)

package react

import (
	"fmt"
	"strings"
)

// writeApiClientEntry emits a single api.<OperationID>(args) entry.
//
// Path params are lifted out of `args` into the openapi-fetch `path` option as
// a typed object literal keyed by the OpenAPI parameter name (e.g.
// `{ contractId: args.contractId }`). Because `args` is typed as Req<K>, a
// wrong key (e.g. args.contractID) is a compile-time error — this statically
// blocks the path-key casing class of defect (BUG-109).
//
// Remaining properties flow as query (GET) or body (mutation). Those records are
// built at runtime and carry a narrow `as any` value cast: the flat Req<K> arg
// shape cannot be statically split into openapi-fetch's per-channel query/body
// types (query may be `never`), so only the query/body *value* is cast — never
// the path object or the whole init. This preserves the static path-key check.
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
		// Typed path object: each key is read from `args` by its OpenAPI name,
		// so a casing mismatch fails tsc. Then a loose view drives the runtime
		// query/body split below.
		var pairs []string
		for _, pp := range ep.pathParams {
			pairs = append(pairs, fmt.Sprintf("%s: args.%s", pp, pp))
		}
		b.WriteString(fmt.Sprintf("    const path = { %s }\n", strings.Join(pairs, ", ")))
		b.WriteString("    const a = args as Record<string, any>\n")
	}

	if method == "GET" {
		if len(ep.pathParams) > 0 {
			b.WriteString("    const query: Record<string, any> = {}\n")
			b.WriteString("    for (const [k, v] of Object.entries(a)) {\n")
			b.WriteString("      if (v == null) continue\n")
			b.WriteString("      if (!(k in path)) query[k] = v\n")
			b.WriteString("    }\n")
			b.WriteString(fmt.Sprintf("    return client.GET('%s', { params: { path, query: query as any } }).then(r => r.data as Res<%s>)\n", pathLit, opQ))
		} else {
			b.WriteString(fmt.Sprintf("    return client.GET('%s', { params: { query: (args ?? {}) as any } }).then(r => r.data as Res<%s>)\n", pathLit, opQ))
		}
	} else {
		verbCall := "POST"
		switch method {
		case "PUT", "PATCH", "DELETE":
			verbCall = method
		}
		if len(ep.pathParams) > 0 {
			b.WriteString("    const body: Record<string, any> = {}\n")
			b.WriteString("    for (const [k, v] of Object.entries(a)) {\n")
			b.WriteString("      if (!(k in path)) body[k] = v\n")
			b.WriteString("    }\n")
			b.WriteString(fmt.Sprintf("    return client.%s('%s', { params: { path }, body: body as any }).then(r => r.data as Res<%s>)\n", verbCall, pathLit, opQ))
		} else {
			b.WriteString(fmt.Sprintf("    return client.%s('%s', { body: (args ?? {}) as any }).then(r => r.data as Res<%s>)\n", verbCall, pathLit, opQ))
		}
	}
	b.WriteString("  },\n")
}
