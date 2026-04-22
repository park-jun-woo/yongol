//ff:func feature=validate type=util control=iteration dimension=1 topic=config-check
//ff:what middlewareSet — manifest middleware 배열을 set으로 변환

package openapi_manifest

func middlewareSet(middleware []string) map[string]bool {
	set := make(map[string]bool, len(middleware))
	for _, m := range middleware {
		set[m] = true
	}
	return set
}
