//ff:func feature=validate type=util control=iteration dimension=1 topic=hurl-openapi
//ff:what joinCSV — 간단 comma-separated join (strings.Join 회피용 로컬 헬퍼)

package hurl_openapi

// joinCSV is a trivial comma-separated join, kept local so we do not
// reach for strings.Join where it would pull in the `strings` import
// only for this one-line use.
func joinCSV(s []string) string {
	out := ""
	for i, v := range s {
		if i > 0 {
			out += ", "
		}
		out += v
	}
	return out
}
