//ff:func feature=validate type=util control=sequence topic=manifest-infra
//ff:what normalizeDDLHead — DDL RawType 에서 배열/파라미터 제거 후 대문자 head 반환

package manifest_ddl

import "strings"

func normalizeDDLHead(rawType string) string {
	t := strings.TrimSpace(rawType)
	if strings.HasSuffix(t, "[]") {
		t = strings.TrimSpace(strings.TrimSuffix(t, "[]"))
	}
	if idx := strings.Index(t, "("); idx > 0 {
		t = t[:idx]
	}
	return strings.ToUpper(strings.TrimSpace(t))
}
