//ff:func feature=report type=test control=iteration dimension=1 topic=sarif
//ff:what test: Emit — SARIF 2.1.0 구조(버전, tool.driver, results, fired rules, level 매핑) 검증
package sarif

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/validate"
)

// TestEmitEmpty covers the minimal empty report — no diagnostics, no rules.
func TestEmitEmpty(t *testing.T) {
	r := &validate.Report{Steps: []validate.StepResult{
		{Name: "step-a", Status: validate.StatusPass},
	}}
	data, err := Emit(r, "v0.1.20", "", nil)
	if err != nil {
		t.Fatalf("Emit: %v", err)
	}
	var doc Document
	if err := json.Unmarshal(data, &doc); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if doc.Version != "2.1.0" {
		t.Errorf("version: got %q, want 2.1.0", doc.Version)
	}
	if doc.Schema == "" {
		t.Errorf("expected $schema to be populated")
	}
	if len(doc.Runs) != 1 {
		t.Fatalf("runs: got %d, want 1", len(doc.Runs))
	}
	run := doc.Runs[0]
	if run.Tool.Driver.Name != "yongol" {
		t.Errorf("driver name: got %q, want yongol", run.Tool.Driver.Name)
	}
	if run.Tool.Driver.Version != "v0.1.20" {
		t.Errorf("driver version: got %q, want v0.1.20", run.Tool.Driver.Version)
	}
	if run.Tool.Driver.InformationURI == "" {
		t.Errorf("informationUri should be populated")
	}
	if len(run.Tool.Driver.Rules) != 0 {
		t.Errorf("rules should be empty for no diagnostics, got %d", len(run.Tool.Driver.Rules))
	}
	if len(run.Results) != 0 {
		t.Errorf("results should be empty, got %d", len(run.Results))
	}
}

// TestEmitWithRuleID verifies rule_id extraction, level mapping and region.
func TestEmitWithRuleID(t *testing.T) {
	r := &validate.Report{Steps: []validate.StepResult{
		{
			Name:   "ssac",
			Status: validate.StatusFail,
			Diagnostics: []diagnostic.Diagnostic{
				{
					File:    "service/auth/login.ssac",
					Line:    15,
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelError,
					Message: "[S-27] 변수 foo 미선언",
				},
				{
					File:    "service/workflow/update.ssac",
					Line:    8,
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelWarning,
					Message: "[S-36] stale @response",
				},
			},
		},
	}}
	data, err := Emit(r, "v0.1.20", "", nil)
	if err != nil {
		t.Fatalf("Emit: %v", err)
	}
	var doc Document
	if err := json.Unmarshal(data, &doc); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	run := doc.Runs[0]
	if len(run.Results) != 2 {
		t.Fatalf("results: got %d, want 2", len(run.Results))
	}
	r0 := run.Results[0]
	if r0.RuleID != "S-27" {
		t.Errorf("results[0].ruleId: got %q, want S-27", r0.RuleID)
	}
	if r0.Level != "error" {
		t.Errorf("results[0].level: got %q, want error", r0.Level)
	}
	if strings.Contains(r0.Message.Text, "[S-27]") {
		t.Errorf("results[0].message.text should not contain [S-27] prefix: %q", r0.Message.Text)
	}
	if len(r0.Locations) != 1 {
		t.Fatalf("results[0].locations: got %d, want 1", len(r0.Locations))
	}
	if r0.Locations[0].PhysicalLocation.Region == nil ||
		r0.Locations[0].PhysicalLocation.Region.StartLine != 15 {
		t.Errorf("results[0].region.startLine: want 15, got %+v",
			r0.Locations[0].PhysicalLocation.Region)
	}

	r1 := run.Results[1]
	if r1.Level != "warning" {
		t.Errorf("results[1].level: got %q, want warning", r1.Level)
	}
	if r1.RuleID != "S-36" {
		t.Errorf("results[1].ruleId: got %q, want S-36", r1.RuleID)
	}

	// tool.driver.rules must contain only rules that actually fired (2 here).
	if len(run.Tool.Driver.Rules) != 2 {
		t.Fatalf("driver.rules: got %d, want 2", len(run.Tool.Driver.Rules))
	}
	wantIDs := map[string]bool{"S-27": false, "S-36": false}
	for _, rl := range run.Tool.Driver.Rules {
		if _, ok := wantIDs[rl.ID]; !ok {
			t.Errorf("unexpected rule id %q", rl.ID)
			continue
		}
		wantIDs[rl.ID] = true
	}
	for id, seen := range wantIDs {
		if !seen {
			t.Errorf("missing rule id %q", id)
		}
	}
}

// TestEmitRuleIDMissing — diagnostic messages without a [RULE-ID] prefix keep
// ruleId empty and do not contribute to tool.driver.rules.
func TestEmitRuleIDMissing(t *testing.T) {
	r := &validate.Report{Steps: []validate.StepResult{
		{
			Name:   "misc",
			Status: validate.StatusFail,
			Diagnostics: []diagnostic.Diagnostic{
				{
					File:    "x.ssac",
					Line:    1,
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelError,
					Message: "plain error message without rule id",
				},
			},
		},
	}}
	data, err := Emit(r, "v0.1.20", "", nil)
	if err != nil {
		t.Fatalf("Emit: %v", err)
	}
	var doc Document
	if err := json.Unmarshal(data, &doc); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	run := doc.Runs[0]
	if len(run.Results) != 1 {
		t.Fatalf("results: got %d, want 1", len(run.Results))
	}
	if run.Results[0].RuleID != "" {
		t.Errorf("ruleId should be empty, got %q", run.Results[0].RuleID)
	}
	if len(run.Tool.Driver.Rules) != 0 {
		t.Errorf("driver.rules should be empty (no fired rule id), got %d",
			len(run.Tool.Driver.Rules))
	}
}
