//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestRenderRefArrayResponseField_ZeroCov — required/optional 분기
package ssac

import (
	"testing"
)

func TestRenderRefArrayResponseField_ZeroCov(t *testing.T) {
	listLocal := map[string]string{"items": "itemsList"}

	req := renderRefArrayResponseField("Items", "items", responseField{IsRequired: true}, listLocal)
	if req != "\tItems: itemsList," {
		t.Errorf("required: got %q", req)
	}

	opt := renderRefArrayResponseField("Items", "items", responseField{IsRequired: false}, listLocal)
	if opt != "\tItems: ptrOf(itemsList)," {
		t.Errorf("optional: got %q", opt)
	}
}
