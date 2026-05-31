//ff:func feature=ssacmeta type=test control=sequence
//ff:what TestloadPackageInterfaceEntry — loadPackageInterfaceEntry() dir/파일/키 폴백 분기
package ssacmeta

import (
	"os"
	"testing"
)

func mustWrite(t *testing.T, path, body string) {
	t.Helper()
	if err := os.WriteFile(path, []byte(body), 0o644); err != nil {
		t.Fatalf("write %s: %v", path, err)
	}
}
