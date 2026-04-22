//ff:func feature=ssac-parse type=parser control=iteration dimension=1
//ff:what 시퀀스에서 @subscribe를 분리하여 ServiceFunc에 할당
package ssac

// filterSubscribe는 시퀀스에서 @subscribe를 분리하여 ServiceFunc에 할당하고 나머지를 반환한다.
func filterSubscribe(sf *ServiceFunc, sequences []Sequence) []Sequence {
	var filtered []Sequence
	for _, seq := range sequences {
		if seq.Type == "subscribe" {
			sf.Subscribe = buildSubscribeInfo(seq.Topic, sf.Param)
			continue
		}
		filtered = append(filtered, seq)
	}
	return filtered
}
