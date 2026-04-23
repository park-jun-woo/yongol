//ff:func feature=gen-gogin type=generator control=sequence
//ff:what blockFileInit — file.Init (local 또는 s3) 블록

package boot

import "github.com/park-jun-woo/yongol/pkg/generate/prepared"

// blockFileInit produces file storage initialization from a resolved
// File. Callers guard with state.ActiveBackends.File != nil so this
// function never sees an inactive subsystem — no raw manifest deref
// possible by signature.
func blockFileInit(f prepared.File) MainBlock {
	backend := f.Backend
	var lines []string
	if backend == "s3" {
		lines = []string{
			`// S3 client setup — requires AWS SDK configuration`,
			`file.Init(file.NewS3File(s3Client, os.Getenv("S3_BUCKET")))`,
		}
	} else {
		lines = []string{`file.Init(file.NewLocalFile(os.Getenv("FILE_ROOT")))`}
	}
	return MainBlock{
		Name: "file-init",
		// Active left nil: collectActiveBlocks appends this block only
		// when prepared.State.ActiveBackends.File != nil.
		Imports: []string{`"github.com/park-jun-woo/ssac/pkg/file"`},
		Lines:   lines,
	}
}
