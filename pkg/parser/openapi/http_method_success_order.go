//ff:func feature=openapi-parse type=model control=sequence
//ff:what httpMethodSuccessOrder — HTTP method 별 2xx 응답 선호 순위 테이블

package openapi

// httpMethodSuccessOrder lists the 2xx status codes yongol considers for a
// given HTTP method, in priority order. The generator emits a response for
// the first entry that the operation actually declares. Values follow RFC
// 9110 REST conventions:
//
//   - POST   : 201 (Created) → 200 (OK)
//   - PUT    : 200 (OK) → 204 (No Content)
//   - PATCH  : 200 (OK) → 204 (No Content)
//   - DELETE : 204 (No Content) → 200 (OK)
//   - GET    : 200 (OK)
//
// SSaC `@post/@put/...` directives describe the DB operation semantics and
// intentionally do not determine HTTP status — HTTP method is the source
// of truth for transport-layer status. Authors may declare multiple 2xx
// responses; yongol picks exactly one.
var httpMethodSuccessOrder = map[string][]int{
	"POST":   {201, 200},
	"PUT":    {200, 204},
	"PATCH":  {200, 204},
	"DELETE": {204, 200},
	"GET":    {200},
}
