//ff:func feature=gen-gogin type=test control=sequence
//ff:what blockFileInit — file.Init (local 또는 s3) 블록
package boot

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
)

func TestBlockFileInit_Local(t *testing.T) {
	block := blockFileInit(prepared.File{Backend: "local"})
	body := strings.Join(block.Lines, "\n")
	if !strings.Contains(body, `file.Init(file.NewLocalFile(os.Getenv("FILE_ROOT")))`) {
		t.Errorf("local backend must use NewLocalFile, got:\n%s", body)
	}
}
