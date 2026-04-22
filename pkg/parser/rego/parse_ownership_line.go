//ff:func feature=policy type=parser control=sequence
//ff:what parseOwnershipLine — 한 줄에서 @ownership 어노테이션 파싱
package rego

func parseOwnershipLine(line string) (OwnershipMapping, bool) {
	m := reOwnership.FindStringSubmatch(line)
	if m == nil {
		return OwnershipMapping{}, false
	}
	om := OwnershipMapping{Resource: m[1], Table: m[2], Column: m[3]}
	if m[4] != "" {
		om.JoinTable = m[4]
		om.JoinFK = m[5]
	}
	return om, true
}
