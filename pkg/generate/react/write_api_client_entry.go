//ff:func feature=gen-react type=generator control=sequence
//ff:what writeApiClientEntry — api.<OperationID>(args) 엔트리 1건 방출 (path/query/body 분기 + r.error throw 승격)

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
//
// withRetry (Phase004, bearer+refresh) wraps the client call in
// withAuthRetry(() => ...) so a 401 triggers the single-flight refresh and
// re-invokes the closure once — the request is rebuilt from args, which is
// why retry lives here and not in the openapi-fetch middleware.
//
// Error promotion (BUG-113): openapi-fetch resolves 4xx/5xx as
// { data: undefined, error: <body> } instead of rejecting. Every entry's
// .then promotes that resolve-with-error to a throw, so TanStack Query's
// onError/error fires with the typed server ErrorResponse and onSuccess
// never sees undefined. The .then sits *outside* wrapAuthRetry: the
// 401→refresh→retry flow completes first and only the final result is
// inspected. Network-level rejects propagate unchanged.
//
// `r.data`/`r.error` are read into `d`/`e` *before* the error check, for
// two tsc reasons. (a) The check would narrow the openapi-fetch result
// union to its success branch, dropping the `| undefined` arm from
// `r.data`; `d` freezes the pre-narrowing type so the `d as Res<K>` cast
// (a `T | undefined → T` widening) stays legal. Res<K> maps the 200 or
// 201 success body (Phase039); only 204/no-body ops keep `void`. (b) The
// load-bearing reason `d`/`e` are pre-captured: for ops declaring no error
// responses the error branch's `error` is `never`, so the check narrows
// `r` itself to `never` and `throw r.error` would not compile; throwing
// the captured `e` stays legal for every op shape.
func writeApiClientEntry(b *strings.Builder, ep endpoint, withRetry bool) {
	method := strings.ToUpper(ep.method)
	pathLit := ep.path
	opQ := fmt.Sprintf("'%s'", ep.opID) // quoted operationId for type arg
	// Shared resolve handler: reject on error, otherwise unwrap typed data.
	thenClause := fmt.Sprintf(".then(r => { const d = r.data; const e = r.error; if (e !== undefined) throw e; return d as Res<%s> })", opQ)

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
			call := wrapAuthRetry(fmt.Sprintf("client.GET('%s', { params: { path, query: query as any } })", pathLit), withRetry)
			b.WriteString(fmt.Sprintf("    return %s%s\n", call, thenClause))
		} else {
			call := wrapAuthRetry(fmt.Sprintf("client.GET('%s', { params: { query: (args ?? {}) as any } })", pathLit), withRetry)
			b.WriteString(fmt.Sprintf("    return %s%s\n", call, thenClause))
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
			call := wrapAuthRetry(fmt.Sprintf("client.%s('%s', { params: { path }, body: body as any })", verbCall, pathLit), withRetry)
			b.WriteString(fmt.Sprintf("    return %s%s\n", call, thenClause))
		} else {
			call := wrapAuthRetry(fmt.Sprintf("client.%s('%s', { body: (args ?? {}) as any })", verbCall, pathLit), withRetry)
			b.WriteString(fmt.Sprintf("    return %s%s\n", call, thenClause))
		}
	}
	b.WriteString("  },\n")
}
