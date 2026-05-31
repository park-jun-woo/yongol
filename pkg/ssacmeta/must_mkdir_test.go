//ff:func feature=ssacmeta type=test control=sequence
//ff:what TestloadPackageInterfaceEntry — loadPackageInterfaceEntry() dir/파일/키 폴백 분기
package ssacmeta

import (
	"os"
	"testing"
)

func mustMkdir(t *testing.T, path string) {
	t.Helper()
	if err := os.MkdirAll(path, 0o755); err != nil {
		t.Fatalf("mkdir %s: %v", path, err)
	}
}
