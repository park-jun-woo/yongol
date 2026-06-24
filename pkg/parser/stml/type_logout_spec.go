//ff:type feature=stml-parse type=model
//ff:what LogoutSpec — 레이아웃 data-logout 속성에서 추출한 세션 종료 선언 구조체

package stml

// LogoutSpec represents a data-logout declaration on a clickable element
// inside a layout (page-flow Phase010). The attribute marks the element
// that ends the session: a value names the server-side session-ending
// operation (e.g. "Logout"), a valueless data-logout declares a
// client-only logout (bearer mode clears the store; cookie mode cannot
// end the session client-side — TM-38 warns).
type LogoutSpec struct {
	OperationID    string // optional OpenAPI operationId of the server logout op ("" when valueless)
	Label          string // element text content (button label)
	RefreshBodyKey string // requestBody property mapped to store.refresh ("" when body-less)
}
