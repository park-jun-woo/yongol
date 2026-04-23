//ff:func feature=report type=util control=sequence topic=sarif
//ff:what regionOrNil — 양수 line 이면 *Region 반환, 아니면 nil
package sarif

// regionOrNil returns a *Region for a positive line number, nil otherwise.
func regionOrNil(line int) *Region {
	if line <= 0 {
		return nil
	}
	return &Region{StartLine: line}
}
