//ff:type feature=stml-parse type=model
//ff:what CaptureBind — data-capture의 응답 필드 → auth sink 바인딩 한 건
package stml

// CaptureBind is a single binding of a data-capture attribute: on action
// success the named response field is stored into the named auth sink
// (e.g. "access_token -> auth.token").
type CaptureBind struct {
	RespField string // OpenAPI 2xx response property name (e.g. "access_token")
	Sink      string // "auth.token" | "auth.refresh" (the only allowed sinks)
}
