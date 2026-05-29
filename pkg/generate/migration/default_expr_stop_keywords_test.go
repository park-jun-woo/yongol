//ff:func feature=migration type=test control=sequence
//ff:what TestDefaultExprStopKeywords — DEFAULT 표현식 종료 키워드 집합 검증
package migration

import "testing"

func TestDefaultExprStopKeywords(t *testing.T) {
	kw := defaultExprStopKeywords()
	for _, want := range []string{"NOT", "NULL", "UNIQUE", "PRIMARY", "REFERENCES", "CHECK", "DEFAULT", "CONSTRAINT", "GENERATED"} {
		if !kw[want] {
			t.Errorf("expected stop keyword %q present", want)
		}
	}
	if kw["SELECT"] {
		t.Error("SELECT should not be a stop keyword")
	}
}
