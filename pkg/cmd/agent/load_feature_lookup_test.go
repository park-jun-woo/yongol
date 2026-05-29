//ff:func feature=agent type=test control=iteration dimension=1
//ff:what TestLoadFeatureLookup — features.yaml에서 op→Feature 맵 구성, 부재 시 nil 검증

package agent

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadFeatureLookup(t *testing.T) {
	dir := t.TempDir()
	yaml := `features:
  - op: Login
    path: /auth/login
    desc: log in
    table: users
  - op: ListOrgs
    path: /orgs
    desc: list orgs
    table: orgs
`
	if err := os.WriteFile(filepath.Join(dir, "features.yaml"), []byte(yaml), 0o644); err != nil {
		t.Fatal(err)
	}

	m := loadFeatureLookup(dir)
	if len(m) != 2 {
		t.Fatalf("lookup len = %d, want 2", len(m))
	}
	if m["Login"].Path != "/auth/login" || m["Login"].Table != "users" {
		t.Errorf("Login = %+v", m["Login"])
	}
	if m["ListOrgs"].Desc != "list orgs" {
		t.Errorf("ListOrgs = %+v", m["ListOrgs"])
	}

	// Missing file → nil.
	if got := loadFeatureLookup(t.TempDir()); got != nil {
		t.Errorf("missing file = %v, want nil", got)
	}
}
