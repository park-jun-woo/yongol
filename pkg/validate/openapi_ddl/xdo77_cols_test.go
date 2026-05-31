//ff:func feature=validate type=test-helper control=sequence topic=openapi-ddl
//ff:what xdo77Cols — XDO-77 테스트용 컬럼명→Go타입 맵 생성 헬퍼
package openapi_ddl

// xdo77Cols returns the shared column→Go-type map for XDO-77 tests.
func xdo77Cols() map[string]string {
	return map[string]string{
		"id":      "int64",
		"email":   "string",
		"active":  "bool",
		"balance": "float64",
	}
}
