//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestNestifyPath_ZeroCov — {param} → :param 변환
package nestjs

import (
	"testing"
)

func TestNestLcFirst_ZeroCov(t *testing.T) {
	if nestLcFirst("") != "" {
		t.Error("empty case")
	}
	if nestLcFirst("Charge") != "charge" {
		t.Error("Charge case")
	}
}
