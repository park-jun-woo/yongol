//ff:func feature=validate type=util control=selection topic=hurl-openapi
//ff:what xoh13IsGuardType — 시퀀스 타입이 guard 인지 판정

package hurl_openapi

func xoh13IsGuardType(seqType string) bool {
	switch seqType {
	case "empty", "exists", "auth", "state", "eval":
		return true
	default:
		return false
	}
}
