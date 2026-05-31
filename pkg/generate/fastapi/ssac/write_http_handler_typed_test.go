//ff:func feature=gen-fastapi type=test control=iteration dimension=1
//ff:what TestWriteHTTPHandlerTyped — 서브테스트 디스패치
package ssac

import "testing"

func TestWriteHTTPHandlerTyped(t *testing.T) {
	for _, st := range []struct {
		name string
		fn   func(*testing.T)
	}{
		{"GETOnlyPath", subtestTestWriteHTTPHandlerTypedGETOnlyPath},
		{"GETWithQuery", subtestTestWriteHTTPHandlerTypedGETWithQuery},
		{"POSTWithBody", subtestTestWriteHTTPHandlerTypedPOSTWithBody},
		{"PUTWithPathAndBody", subtestTestWriteHTTPHandlerTypedPUTWithPathAndBody},
		{"DELETENoBody", subtestTestWriteHTTPHandlerTypedDELETENoBody},
	} {
		t.Run(st.name, st.fn)
	}
}
