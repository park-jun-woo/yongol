//ff:func feature=rule type=test-helper control=sequence
//ff:what runBaseSpecValidateCase — TestBaseSpec_Validate 단일 테이블 케이스용 서브테스트 클로저 팩토리

package rule

import "testing"

func runBaseSpecValidateCase(spec BaseSpec, wantErr bool, wantSub string) func(t *testing.T) {
	return func(t *testing.T) {
		err := spec.Validate()
		if wantErr {
			checkBaseSpecValidateExpectedError(t, err, wantSub)
			return
		}
		if err != nil {
			t.Fatalf("Validate() = %v; want nil", err)
		}
	}
}
