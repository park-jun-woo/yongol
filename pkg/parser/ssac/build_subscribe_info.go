//ff:func feature=ssac-parse type=parser control=sequence
//ff:what SubscribeInfo를 생성
package ssac

// buildSubscribeInfo는 SubscribeInfo를 생성한다.
func buildSubscribeInfo(topic string, param *ParamInfo) *SubscribeInfo {
	si := &SubscribeInfo{Topic: topic}
	if param != nil {
		si.MessageType = param.TypeName
	}
	return si
}
