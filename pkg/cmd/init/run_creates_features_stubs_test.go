//ff:func feature=cli-init type=test control=iteration dimension=1
//ff:what TestRunCreatesFeaturesStubs

package cliinit

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunCreatesFeaturesStubs(t *testing.T) {
	featPath := writeTempFeatures(t)
	tmp := t.TempDir()
	target := filepath.Join(tmp, "myapp")
	var outBuf, errBuf bytes.Buffer
	err := Run(&outBuf, &errBuf, Options{
		ProjectID:    "Myapp",
		FeaturesPath: featPath,
		Description:  "Test features project",
		Dir:          target,
		Module:       "github.com/test/myapp",
	})
	if err != nil {
		t.Fatalf("Run returned error: %v", err)
	}

	// OpenAPI should contain operationIds from features.
	openapi, err := os.ReadFile(filepath.Join(target, "specs", "api", "openapi.yaml"))
	if err != nil {
		t.Fatalf("read openapi: %v", err)
	}
	openapiStr := string(openapi)
	for _, op := range []string{"CreateWorkflow", "GetWorkflow", "ListWorkflows"} {
		if !strings.Contains(openapiStr, "operationId: "+op) {
			t.Errorf("openapi missing operationId: %s", op)
		}
	}

	// SSaC stub files should exist per feature.
	for _, op := range []string{"CreateWorkflow", "GetWorkflow", "ListWorkflows"} {
		ssacPath := filepath.Join(target, "specs", "service", "workflow", op+".ssac")
		if _, err := os.Stat(ssacPath); err != nil {
			t.Errorf("expected SSaC file %s missing: %v", ssacPath, err)
		}
	}

	// Rego should contain allow rules for each feature.
	rego, err := os.ReadFile(filepath.Join(target, "specs", "policy", "authz.rego"))
	if err != nil {
		t.Fatalf("read rego: %v", err)
	}
	regoStr := string(rego)
	if !strings.Contains(regoStr, `"CreateWorkflow"`) {
		t.Errorf("rego missing CreateWorkflow allow rule")
	}
	if !strings.Contains(regoStr, `"workflow"`) {
		t.Errorf("rego missing workflow resource")
	}

	// Hurl smoke test should exist.
	hurlContent, err := os.ReadFile(filepath.Join(target, "specs", "tests", "smoke.hurl"))
	if err != nil {
		t.Fatalf("read hurl: %v", err)
	}
	hurlStr := string(hurlContent)
	if !strings.Contains(hurlStr, "# CreateWorkflow") {
		t.Errorf("hurl missing CreateWorkflow request")
	}
	if !strings.Contains(hurlStr, "POST {{host}}/workflows") {
		t.Errorf("hurl missing POST /workflows")
	}

	// features.yaml should be copied.
	if _, err := os.Stat(filepath.Join(target, "specs", "features.yaml")); err != nil {
		t.Errorf("features.yaml not copied: %v", err)
	}

	// .yongol hash file should exist.
	yongolData, err := os.ReadFile(filepath.Join(target, "specs", ".yongol"))
	if err != nil {
		t.Fatalf("read .yongol: %v", err)
	}
	yongolStr := string(yongolData)
	if !strings.Contains(yongolStr, "sha256:") {
		t.Errorf(".yongol missing sha256 hash")
	}

	// Output should mention feature count.
	if !strings.Contains(outBuf.String(), "3 features") {
		t.Errorf("output should mention feature count, got: %s", outBuf.String())
	}
}
