//ff:func feature=rule type=test control=sequence topic=catalog
//ff:what zz_zerocov_test — catalog.Index / catalog.MustLoad 0% 커버리지 단위 테스트
package catalog

import (
	"testing"
)

func TestMustLoad_ZeroCov(t *testing.T) {
	// MustLoad on the embedded catalog must succeed (no log.Fatal path).
	c := MustLoad()
	if c == nil {
		t.Fatal("MustLoad returned nil")
	}
}
