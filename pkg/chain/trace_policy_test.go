//ff:func feature=chain type=test control=iteration dimension=3
//ff:what tracePolicy 가 @auth resource 를 Rego rule 과 매칭하고 action 필터링/중복제거/미참조 nil 을 처리하는지 검증
package chain

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/rego"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestTracePolicy(t *testing.T) {
	specsDir := t.TempDir()
	authzDir := filepath.Join(specsDir, "authz")
	if err := os.MkdirAll(authzDir, 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	regoFile := filepath.Join(authzDir, "project.rego")
	content := "package project\n\nallow if input.resource == \"project\"\n"
	if err := os.WriteFile(regoFile, []byte(content), 0o644); err != nil {
		t.Fatalf("write: %v", err)
	}

	policies := []rego.Policy{
		{
			File: regoFile,
			Rules: []rego.AllowRule{
				{Resource: "project", Actions: []string{"delete", "update"}},
				{Resource: "other", Actions: []string{"read"}}, // not referenced
			},
		},
	}

	sf := &ssac.ServiceFunc{
		Name: "DeleteProject",
		Sequences: []ssac.Sequence{
			{Type: "auth", Resource: "project", Action: "delete"},
			{Type: "get", Model: "Project.FindByID"}, // ignored
		},
	}

	links := tracePolicy(sf, policies, specsDir)
	if len(links) != 1 {
		t.Fatalf("expected 1 rego link, got %d: %+v", len(links), links)
	}
	l := links[0]
	if l.Kind != "Rego" || l.File != "authz/project.rego" {
		t.Errorf("link fields: %+v", l)
	}
	// grepLine matches the first occurrence of the resource name ("project"),
	// which is the package declaration on line 1.
	if l.Line != 1 {
		t.Errorf("line: got %d, want 1", l.Line)
	}
	// Only the requested action "delete" should be listed, not "update".
	if l.Summary != "resource: project [delete]" {
		t.Errorf("summary: got %q, want %q", l.Summary, "resource: project [delete]")
	}

	// No @auth resources → nil.
	sfNone := &ssac.ServiceFunc{Name: "X", Sequences: []ssac.Sequence{{Type: "get", Model: "Y.Z"}}}
	if tracePolicy(sfNone, policies, specsDir) != nil {
		t.Error("expected nil when no @auth resources")
	}

	// Dedup: two rules in the same file both match the referenced resource;
	// the second hits the seen[relPath] guard so only one link is produced.
	t.Run("Dedup", func(t *testing.T) {
		dupPolicies := []rego.Policy{
			{
				File: regoFile,
				Rules: []rego.AllowRule{
					{Resource: "project", Actions: []string{"delete"}},
					{Resource: "project", Actions: []string{"update"}}, // same file+resource
				},
			},
		}
		links := tracePolicy(sf, dupPolicies, specsDir)
		if len(links) != 1 {
			t.Fatalf("expected 1 deduped rego link, got %d: %+v", len(links), links)
		}
		if links[0].File != "authz/project.rego" {
			t.Errorf("link file: got %q, want authz/project.rego", links[0].File)
		}
	})

	// NoActionMatch: the @auth resource matches a rule, but none of the rule's
	// actions are requested by the SSaC function (and the auth seq carries an
	// empty Action), so actList stays empty and the summary has no bracket
	// suffix. Exercises the `seq.Action == ""` skip and `len(actList) == 0`
	// branch of the summary builder.
	t.Run("NoActionMatch", func(t *testing.T) {
		sfNoAction := &ssac.ServiceFunc{
			Name: "ListProject",
			Sequences: []ssac.Sequence{
				{Type: "auth", Resource: "project", Action: ""}, // empty action
			},
		}
		pol := []rego.Policy{
			{
				File: regoFile,
				Rules: []rego.AllowRule{
					{Resource: "project", Actions: []string{"delete", "update"}},
				},
			},
		}
		links := tracePolicy(sfNoAction, pol, specsDir)
		if len(links) != 1 {
			t.Fatalf("expected 1 link, got %d: %+v", len(links), links)
		}
		if links[0].Summary != "resource: project" {
			t.Errorf("summary: got %q, want %q (no action brackets)", links[0].Summary, "resource: project")
		}
	})
}
