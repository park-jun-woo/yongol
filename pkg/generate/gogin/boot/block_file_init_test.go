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

func TestBlockFileInit_S3(t *testing.T) {
	block := blockFileInit(prepared.File{Backend: "s3"})
	body := strings.Join(block.Lines, "\n")
	if !strings.Contains(body, `file.Init(file.NewS3File(s3Client, os.Getenv("S3_BUCKET")))`) {
		t.Errorf("s3 backend must use NewS3File, got:\n%s", body)
	}
	if !strings.Contains(strings.Join(block.Imports, "\n"), `"github.com/park-jun-woo/ssac/pkg/file"`) {
		t.Errorf("must import ssac file pkg, got:\n%v", block.Imports)
	}
}
