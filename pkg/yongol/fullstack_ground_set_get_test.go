//ff:func feature=orchestrator type=test control=sequence
//ff:what TestFullstack Ground/SetGround/PresenceOf 및 indexDetected 검증
package yongol

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/rule"
)

func TestFullstackGroundSetGet(t *testing.T) {
	fs := &Fullstack{}
	if fs.Ground() != nil {
		t.Error("expected nil Ground before SetGround")
	}
	g := &rule.Ground{}
	fs.SetGround(g)
	if fs.Ground() != g {
		t.Error("Ground() did not return the bound rule.Ground")
	}
}
