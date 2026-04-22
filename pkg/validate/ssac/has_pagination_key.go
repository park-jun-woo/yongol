//ff:func feature=validate type=util control=iteration dimension=1 topic=ssac-structural
//ff:what hasPaginationKey — Inputs 에 pagination 키 (Page/PerPage/Cursor) 존재 여부

package ssac

// hasPaginationKey returns true when inputs contains at least one pagination key.
func hasPaginationKey(inputs map[string]string) bool {
	for k := range inputs {
		if paginationKeys[k] {
			return true
		}
	}
	return false
}
