//ff:func feature=funcspec type=parser control=iteration dimension=1
//ff:what FuncSpec의 빈 RequestFields/ResponseFields를 패키지 레벨 타입에서 보충하고 수집 에러는 Diagnostic 으로 전파한다
package funcspec

import "github.com/park-jun-woo/yongol/pkg/diagnostic"

// fillMissingFields fills empty RequestFields/ResponseFields from
// companion struct files in the same directory.
//
// collectPackageTypes 가 반환하는 Diagnostic 은 같은 디렉토리에 대한
// 중복 append 를 피하도록 seenDir 로 dedup 한다.
func fillMissingFields(specs []FuncSpec, specDirs []string) []diagnostic.Diagnostic {
	cache := make(map[string]map[string][]Field)
	seenDir := make(map[string]struct{})
	var diags []diagnostic.Diagnostic
	for i := range specs {
		if len(specs[i].RequestFields) > 0 && len(specs[i].ResponseFields) > 0 {
			continue
		}
		dir := specDirs[i]
		typeMap, ok := cache[dir]
		if !ok {
			var d []diagnostic.Diagnostic
			typeMap, d = collectPackageTypes(dir)
			cache[dir] = typeMap
			if _, seen := seenDir[dir]; !seen {
				diags = append(diags, d...)
				seenDir[dir] = struct{}{}
			}
		}
		fillSpecFromTypeMap(&specs[i], typeMap)
	}
	return diags
}
