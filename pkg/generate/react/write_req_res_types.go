//ff:func feature=gen-react type=generator control=sequence
//ff:what writeReqResTypes — PathOf/QueryOf/BodyOf + Req<K> / Res<K> 타입 헬퍼를 방출한다

package react

import "strings"

// writeReqResTypes emits PathOf, QueryOf, BodyOf helper types and the
// Req<K>/Res<K> intersection type that unions path params, query params, and
// request body into a single flat argument type per operation.
func writeReqResTypes(b *strings.Builder) {
	b.WriteString("type PathOf<K extends keyof operations> =\n")
	b.WriteString("  operations[K] extends { parameters: { path: infer P } } ? P : {}\n\n")

	b.WriteString("type QueryOf<K extends keyof operations> =\n")
	b.WriteString("  operations[K] extends { parameters: { query?: infer Q } } ? (Q extends undefined ? {} : Q) : {}\n\n")

	b.WriteString("type BodyOf<K extends keyof operations> =\n")
	b.WriteString("  operations[K] extends { requestBody: { content: { 'application/json': infer B } } }\n")
	b.WriteString("    ? B extends Record<string, never> ? {} : B\n")
	b.WriteString("  : operations[K] extends { requestBody: { content: { 'multipart/form-data': infer B } } }\n")
	b.WriteString("    ? B extends Record<string, never> ? {} : B\n")
	b.WriteString("    : {}\n\n")

	b.WriteString("type Req<K extends keyof operations> =\n")
	b.WriteString("  PathOf<K> & QueryOf<K> & BodyOf<K> extends infer R\n")
	b.WriteString("    ? keyof R extends never ? void : R\n")
	b.WriteString("    : void\n\n")

	b.WriteString("type Res<K extends keyof operations> =\n")
	b.WriteString("  operations[K] extends { responses: { 200: { content: { 'application/json': infer R } } } } ? R : void\n\n")
}
