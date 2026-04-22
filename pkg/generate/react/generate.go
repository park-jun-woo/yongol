//ff:func feature=gen-react type=generator control=sequence
//ff:what Generate — Fullstack → React + Vite + TanStack Query + shadcn-like 스캐폴드 방출

package react

import (
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// Generate materializes the full React scaffold under <artifactsDir>/frontend/.
// Emitted artifacts:
//
//	index.html, package.json, tsconfig.json, vite.config.ts
//	tailwind.config.js, postcss.config.js
//	src/main.tsx, src/App.tsx, src/index.css
//	src/api.ts  — operationId-keyed apiClient (openapi-fetch + typed paths)
//	src/types/api.d.ts — openapi-typescript output (spawn)
//	src/lib/utils.ts — shadcn cn() helper
//	src/components/ui/*.tsx — 10 shadcn-like primitives
//
// Page files (src/pages/*.tsx) are **not** emitted — TSX is the SSOT, so
// yongol reads them, not writes them.
func Generate(fs *yongol.Fullstack, artifactsDir string) error {
	return generateFrontendSetup(fs, artifactsDir)
}
