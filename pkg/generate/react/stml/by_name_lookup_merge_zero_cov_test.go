//ff:func feature=gen-react type=test control=sequence
//ff:what TestByName_ZeroCov — react/stml 코드젠 헬퍼들을 이름으로 직접 호출해 커버리지 귀속
package stml

import (
	"testing"
)

func TestByNameLookupMerge_ZeroCov(t *testing.T) {
	cons := byNameConstraints()
	if lookupConstraints("CreateItem", cons) == nil {
		t.Errorf("lookupConstraints nil")
	}
	_ = lookupConstraints("Missing", cons)
	_ = lookupConstraints("X", nil)

	base := DefaultOptions()
	merged := mergeOpt(base, GenerateOptions{
		APIImportPath:           "@/api",
		UseClient:               true,
		BearerAuth:              true,
		RequestConstraints:      cons,
		ResponseArrayItemFields: map[string]map[string]map[string]bool{"ListItems": {"items": {"id": true}}},
		NoBodyOps:               map[string]bool{"X": true},
		PathParamTypes:          map[string]map[string]string{"GetItem": {"id": "integer"}},
	})
	if merged.APIImportPath != "@/api" {
		t.Errorf("mergeOpt APIImportPath = %q", merged.APIImportPath)
	}
}
