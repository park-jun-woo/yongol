//ff:func feature=cli type=test-helper control=sequence
//ff:what subtestTestPrintReportSARIFCatalogRulesEmbedded — CatalogRulesEmbedded 서브테스트
package main

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/validate"
)

func subtestTestPrintReportSARIFCatalogRulesEmbedded(t *testing.T) {

	// Exercises the rulecatalog.Load + sarif.Emit(cat) path: the embedded
	// rulebook must populate runs[0].tool.driver.rules[] (non-empty), proving
	// the non-error catalog branch wired the catalog into the SARIF document.
	r := &validate.Report{Steps: []validate.StepResult{
		{Name: "step", Status: validate.StatusPass},
	}}
	var buf bytes.Buffer
	if _, _, err := printReportSARIF(&buf, r, "/tmp/specs"); err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	var doc struct {
		Runs []struct {
			Tool struct {
				Driver struct {
					Rules []map[string]any `json:"rules"`
				} `json:"driver"`
			} `json:"tool"`
		} `json:"runs"`
	}
	if err := json.Unmarshal(buf.Bytes(), &doc); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}
	if len(doc.Runs) == 0 {
		t.Fatal("expected at least one run in SARIF document")
	}
	if len(doc.Runs[0].Tool.Driver.Rules) == 0 {
		t.Fatal("expected embedded rule catalog to populate tool.driver.rules[]")
	}

}
