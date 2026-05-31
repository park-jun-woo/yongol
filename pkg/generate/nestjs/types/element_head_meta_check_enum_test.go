//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestElementHeadMetaCheckEnum — elementHeadMeta.CheckEnum nil 반환 커버
package types

import "testing"

func TestElementHeadMetaCheckEnum_ZeroCov(t *testing.T) {
	m := elementHeadMeta{head: "TEXT"}
	if m.CheckEnum() != nil {
		t.Errorf("CheckEnum should be nil")
	}
}
