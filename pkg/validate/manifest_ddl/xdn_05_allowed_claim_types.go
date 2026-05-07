//ff:func feature=validate type=util control=selection topic=manifest-infra
//ff:what allowedClaimTypes — XDN-05 허용 claim 타입 집합

package manifest_ddl

var allowedClaimTypes = map[string]bool{
	"string": true,
	"int64":  true,
	"int32":  true,
	"bool":   true,
	"uuid":   true,
}
