//ff:func feature=gen-react type=test control=sequence
//ff:what TestByName_ZeroCov — react/stml 코드젠 헬퍼들을 이름으로 직접 호출해 커버리지 귀속
package stml

import (
	"testing"
)

func TestByNameZodValidations_ZeroCov(t *testing.T) {
	cons := byNameConstraints()
	strFC := cons["Login"]["Email"]
	numFC := cons["CreateItem"]["Count"]

	parts := appendStringValidations(nil, strFC)
	if len(parts) == 0 {
		t.Errorf("appendStringValidations empty")
	}
	_ = appendStringValidations(nil, numFC) // non-string returns unchanged

	nparts := appendNumericValidations(nil, numFC)
	if len(nparts) == 0 {
		t.Errorf("appendNumericValidations empty")
	}
	_ = appendNumericValidations(nil, strFC) // non-numeric returns unchanged
}
