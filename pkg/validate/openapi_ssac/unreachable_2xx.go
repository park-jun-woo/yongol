//ff:func feature=validate type=util control=iteration dimension=1 topic=ssac-openapi
//ff:what XOS-82 헬퍼 — 선언된 2xx 중 yongol 이 emit 하지 않는 상태 집합 반환

package openapi_ssac

// unreachable2xx returns the subset of declared 2xx status codes that
// yongol will not emit given the currently selected success status.
func unreachable2xx(declared map[int]bool, selected int) map[int]bool {
	out := make(map[int]bool, len(declared))
	for code := range declared {
		if code == selected {
			continue
		}
		out[code] = true
	}
	return out
}
