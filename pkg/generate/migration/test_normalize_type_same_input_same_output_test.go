//ff:func feature=migration type=test control=sequence
//ff:what TestNormalizeType_SameInputSameOutput — int/int4/INTEGER 등 alias 가 동일 CanonicalType 반환
package migration

import "testing"

func TestNormalizeType_SameInputSameOutput(t *testing.T) {
	ct1, _ := NormalizeType("int")
	ct2, _ := NormalizeType("int4")
	ct3, _ := NormalizeType("INTEGER")
	if !ct1.Equal(ct2) || !ct2.Equal(ct3) {
		t.Errorf("INTEGER aliases not equal: %+v %+v %+v", ct1, ct2, ct3)
	}
	ts1, _ := NormalizeType("timestamptz")
	ts2, _ := NormalizeType("timestamp with time zone")
	if !ts1.Equal(ts2) {
		t.Errorf("TIMESTAMPTZ aliases not equal: %+v %+v", ts1, ts2)
	}
}
