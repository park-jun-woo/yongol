//ff:func feature=agent type=test control=iteration dimension=1
//ff:what TestTopoSortTables — belongs_to 위상 정렬(부모 우선)과 순환 시 깨고 반환 검증
package agent

func indexOf(s []string, v string) int {
	for i, x := range s {
		if x == v {
			return i
		}
	}
	return -1
}
