//ff:func feature=external type=util control=iteration dimension=1
//ff:what 구조체 타입 목록에서 time.Time 사용 여부를 확인한다
package external

func typesNeedTime(types []structType) bool {
	for _, t := range types {
		if structHasTimeField(t) {
			return true
		}
	}
	return false
}
