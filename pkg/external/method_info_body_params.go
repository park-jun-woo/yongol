//ff:func feature=external type=generator control=iteration dimension=1
//ff:what 바디 파라미터만 필터링하여 반환한다
package external

func (m methodInfo) bodyParams() []paramInfo {
	var result []paramInfo
	for _, p := range m.Params {
		if p.In == "body" {
			result = append(result, p)
		}
	}
	return result
}
