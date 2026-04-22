//ff:func feature=validate type=accessor control=iteration dimension=1
//ff:what Report.HasFailure — 임의 step이 StatusFail이면 true 반환
package validate

// HasFailure reports whether any step ended in StatusFail.
func (r *Report) HasFailure() bool {
	for _, s := range r.Steps {
		if s.Status == StatusFail {
			return true
		}
	}
	return false
}
