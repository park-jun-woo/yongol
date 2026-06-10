//ff:func feature=gen-react type=generator control=sequence
//ff:what package.json — React 19 + Vite + TanStack Query + shadcn 에코시스템 의존성

package react

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// writePackageJSON emits the frontend package.json. The dependency set is
// deliberate:
//   - react/react-dom 19, react-router-dom 7 (current majors as of 2026-04)
//   - @tanstack/react-query 5 for server state
//   - react-hook-form 7 for form state (XOT-3 depends on its register())
//   - openapi-fetch 0.x — typed fetch client consumed by src/lib/api.ts
//   - openapi-typescript 7.x (devDep) — `postinstall` runs on install; yongol
//     also re-runs it from `yongol generate` so generated docs are fresh.
//   - tailwindcss 3.x + clsx + tailwind-merge — shadcn/ui primitives rely on
//     the cn() utility.
//   - zustand 5 (withAuthStore only) — the bearer session store emitted as
//     src/stores/auth.ts depends on it.
func writePackageJSON(dir string, withAuthStore bool) error {
	deps := map[string]string{
		"react":                 "^19",
		"react-dom":             "^19",
		"react-router-dom":      "^7",
		"@tanstack/react-query": "^5",
		"react-hook-form":       "^7",
		"zod":                   "^3",
		"@hookform/resolvers":   "^3",
		"openapi-fetch":         "^0.13",
		"clsx":                  "^2",
		"tailwind-merge":        "^2",
	}
	if withAuthStore {
		deps["zustand"] = "^5"
	}
	pkg := map[string]interface{}{
		"private": true,
		"type":    "module",
		"scripts": map[string]string{
			"dev":     "vite",
			"build":   "tsc -b && vite build",
			"preview": "vite preview",
			"gen:api": "openapi-typescript ../specs/api/openapi.yaml -o ./src/types/api.d.ts",
		},
		"dependencies": deps,
		"devDependencies": map[string]string{
			"@types/react":         "^19",
			"@types/react-dom":     "^19",
			"@vitejs/plugin-react": "^4",
			"typescript":           "^5",
			"vite":                 "^5",
			"tailwindcss":          "^3",
			"postcss":              "^8",
			"autoprefixer":         "^10",
			"openapi-typescript":   "^7",
		},
	}
	b, err := json.MarshalIndent(pkg, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(dir, "package.json"), b, 0644)
}
