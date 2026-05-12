//ff:type feature=generate type=model
//ff:what STML 폼 액션의 operationId와 필드 이름 목록을 나타내는 구조체
package generate

// actionEntry pairs an operationId with the STML field names from its form.
type actionEntry struct {
	opID       string
	fieldNames []string
}
