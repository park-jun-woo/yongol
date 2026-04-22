//ff:func feature=external type=util control=iteration dimension=1
//ff:what 구조체 필드 중 time.Time 타입이 있는지 확인한다
package external

func structHasTimeField(t structType) bool {
	for _, f := range t.Fields {
		if f.GoType == "time.Time" {
			return true
		}
	}
	return false
}
