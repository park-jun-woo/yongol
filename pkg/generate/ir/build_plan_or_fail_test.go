//ff:func feature=gen-ir type=test control=sequence
//ff:what TestConvertersZeroCov — 각 SSaC 시퀀스 종류를 BuildServicePlan 으로 변환해 convert* 디스패처 전체 커버
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func buildPlanOrFail(t *testing.T, name string, seqs []ssac.Sequence) *ServicePlan {
	t.Helper()
	sf := &ssac.ServiceFunc{
		Name:      name,
		FileName:  name + ".ssac",
		Feature:   "feature",
		Sequences: seqs,
	}
	plan, err := BuildServicePlan(sf, &yongol.Fullstack{})
	if err != nil {
		t.Fatalf("%s: BuildServicePlan error: %v", name, err)
	}
	if plan == nil {
		t.Fatalf("%s: nil plan", name)
	}
	return plan
}
