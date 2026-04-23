//ff:func feature=migration type=test control=iteration dimension=1
//ff:what TestBuildASTFromDir_Zenflow — zenflow 더미 스펙 디렉토리 파싱 결과에 핵심 테이블 존재 확인
package migration

import (
	"path/filepath"
	"runtime"
	"testing"
)

func TestBuildASTFromDir_Zenflow(t *testing.T) {
	_, thisFile, _, _ := runtime.Caller(0)
	root := filepath.Join(filepath.Dir(thisFile), "..", "..", "..")
	ddlDir := filepath.Join(root, "dummys", "zenflow", "try-02", "specs", "db")
	s, err := BuildASTFromDir(ddlDir, []string{SnapshotFileName})
	if err != nil {
		t.Fatalf("BuildASTFromDir: %v", err)
	}
	expect := []string{"users", "organizations", "workflows", "actions"}
	for _, name := range expect {
		if _, ok := s.Tables[name]; !ok {
			t.Errorf("table %q missing from AST (have %v)", name, mapKeys(s.Tables))
		}
	}
}
