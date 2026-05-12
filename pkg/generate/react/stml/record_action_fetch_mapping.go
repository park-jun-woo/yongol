//ff:func feature=stml-gen type=util control=sequence
//ff:what action OperationID와 parent fetch ops의 매핑을 기록한다

package stml

// recordActionFetchMapping records the mapping between an action's OperationID
// and its ancestor fetch block OperationIDs.
func recordActionFetchMapping(opID string, parentFetchOps []string, m map[string][]string) {
	if _, ok := m[opID]; ok {
		return
	}
	if len(parentFetchOps) > 0 {
		m[opID] = append([]string{}, parentFetchOps...)
	}
}
