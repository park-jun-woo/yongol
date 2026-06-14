//ff:func feature=stml-gen type=util control=iteration dimension=1
//ff:what delete 액션과 같은 리소스를 가리키는 자기 GET operationId를 fetchOps에서 찾는다 (이름 짝 + path-param 매칭)
package stml

import "strings"

// selfGetOps returns the page fetch operationId that reads the very item a
// delete action removes — its "self GET" (BUG-132 132-2). The delete's
// resource GET is paired by name (Delete<Entity> -> Get<Entity>) and
// confirmed by an identical path-param signature, so a delete on
// /buildings/{buildingId} matches GetBuilding but not the sibling queries
// CheckBuildingDeletable or ListBuildingPhotos that share the path param.
// After the delete this query is a guaranteed 404, so the caller emits
// removeQueries for it instead of invalidateQueries. Returns nil when the
// delete has no path params or its self GET is not fetched on the page.
func selfGetOps(deleteOpID string, fetchOps []string, pathParamTypes map[string]map[string]string) []string {
	entity := strings.TrimPrefix(deleteOpID, "Delete")
	if entity == deleteOpID || entity == "" {
		return nil
	}
	candidate := "Get" + entity
	delParams := pathParamTypes[deleteOpID]
	if len(delParams) == 0 {
		return nil
	}
	for _, op := range fetchOps {
		if op == candidate && pathParamSetEqual(delParams, pathParamTypes[op]) {
			return []string{candidate}
		}
	}
	return nil
}
