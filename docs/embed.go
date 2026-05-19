package docs

import "embed"

//go:embed ssac.md ddl.md openapi.md policy.md states.md scenario.md manifest.md func.md
var FS embed.FS
