//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what authStoreImports 단위 테스트 (subscribe/wrap_calls/gin 의존 import 조립)
package ssac

func contains(s []string, want string) bool {
	for _, v := range s {
		if v == want {
			return true
		}
	}
	return false
}
