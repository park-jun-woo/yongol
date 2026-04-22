//ff:type feature=validate type=model topic=policy-check
//ff:what PairLocation — (action, resource) 출현 위치 (File + Line)

package ssac_rego

// PairLocation captures the source-file/line of a single occurrence.
type PairLocation struct {
	File string
	Line int
}
