//ff:func feature=gen-gogin type=test control=sequence
//ff:what methodGen.returnErr 단위 테스트 (subscribe면 return err, 아니면 return nil, err)
package ssac

import (
	"testing"
)

func TestMethodGenReturnErr(t *testing.T) {
	if got := (&methodGen{IsSubscribe: true}).returnErr(); got != "return err" {
		t.Errorf("subscribe returnErr = %q, want %q", got, "return err")
	}
	if got := (&methodGen{IsSubscribe: false}).returnErr(); got != "return nil, err" {
		t.Errorf("http returnErr = %q, want %q", got, "return nil, err")
	}
}
