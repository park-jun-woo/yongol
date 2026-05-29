//ff:func feature=gen-ir type=test control=sequence
//ff:what TestBuildInfraPlanEmpty -- 비활성 상태 → InfraPlan 모든 필드 nil 검증

package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBuildInfraPlanEmpty(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &manifest.ProjectConfig{
			Backend: manifest.Backend{Module: "github.com/test/empty"},
		},
	}
	ps := &prepared.State{}

	plan := BuildInfraPlan(fs, ps)

	if plan.Session != nil {
		t.Error("Session should be nil when inactive")
	}
	if plan.Cache != nil {
		t.Error("Cache should be nil when inactive")
	}
	if plan.Queue != nil {
		t.Error("Queue should be nil when inactive")
	}
	if plan.Auth != nil {
		t.Error("Auth should be nil when inactive")
	}
}
