//ff:func feature=gen-filefunc type=util control=iteration dimension=1
//ff:what ensureGenFilefuncEntry — inserts the fixed feature entries that the generator always requires
package filefunc

// anchorFeatures are feature keys always present in the generated codebook
// because specific generator outputs hardcode them on //ff:func headers.
var anchorFeatures = map[string]string{
	"gen-filefunc":       "filefunc codebook.yaml auto-generation",
	"runtime-middleware": "runtime request validation middleware (kin-openapi)",
	"main":               "application entry point (cmd/main.go)",
}

// ensureGenFilefuncEntry guarantees that the codebook contains the anchor
// feature entries so downstream validators accept generator-hardcoded
// feature names even when no on-disk package carries them.
func ensureGenFilefuncEntry(dst map[string]string) {
	for name, desc := range anchorFeatures {
		insertFeatureIfNew(dst, name, desc)
	}
}
