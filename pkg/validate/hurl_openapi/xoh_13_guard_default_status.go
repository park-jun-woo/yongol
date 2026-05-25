//ff:func feature=validate type=util control=selection topic=hurl-openapi
//ff:what xoh13GuardDefaultStatus — guard 타입별 기본 ErrStatus 반환

package hurl_openapi

func xoh13GuardDefaultStatus(seqType string) int {
	switch seqType {
	case "empty":
		return 404
	case "exists":
		return 409
	case "auth":
		return 403
	case "state":
		return 409
	case "eval":
		return 0
	default:
		return 0
	}
}
