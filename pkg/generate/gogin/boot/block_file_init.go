//ff:func feature=gen-gogin type=generator control=sequence
//ff:what blockFileInit — file.Init (local 또는 s3) 블록

package boot

import "github.com/park-jun-woo/yongol/pkg/yongol"

// blockFileInit produces file storage initialization. Active when
// manifest.file.backend is set.
func blockFileInit(fs *yongol.Fullstack) MainBlock {
	backend := fs.Manifest.File.Backend
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
		Name:    "file-init",
		Active:  hasFile,
		Imports: []string{`"github.com/park-jun-woo/ssac/pkg/file"`},
		Lines:   lines,
	}
}
