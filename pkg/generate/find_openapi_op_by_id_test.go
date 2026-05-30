//ff:func feature=generate type=test control=sequence
//ff:what findOpenAPIOpByID의 nil 문서/매칭/미매칭 분기 검증

package generate

import "testing"

func TestFindOpenAPIOpByID(t *testing.T) {
	t.Run("NilDoc", func(t *testing.T) {
		if _, ok := findOpenAPIOpByID(nil, "CreateItem"); ok {
			t.Errorf("expected not found for nil doc")
		}
	})

	t.Run("Found", func(t *testing.T) {
		doc := loadTestDoc(t)
		op, ok := findOpenAPIOpByID(doc, "CreateItem")
		if !ok {
			t.Fatalf("expected CreateItem to be found")
		}
		if op == nil || op.OperationID != "CreateItem" {
			t.Errorf("expected CreateItem operation, got: %v", op)
		}
	})

	t.Run("NotFound", func(t *testing.T) {
		doc := loadTestDoc(t)
		if _, ok := findOpenAPIOpByID(doc, "Nonexistent"); ok {
			t.Errorf("expected not found for unknown operationID")
		}
	})
}
