//ff:func feature=report type=test control=iteration dimension=1 topic=sarif
//ff:what TestEmitWithRuleID — rule_id 추출 + level 매핑 + region + fired rules 검증
package sarif

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/validate"
)

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
					Message: "[S-27] variable foo is not declared",
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
