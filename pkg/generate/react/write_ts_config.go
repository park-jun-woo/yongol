//ff:func feature=gen-react type=generator control=sequence
//ff:what tsconfig.json — @/ path-alias + strict 활성화

package react

import (
	"os"
	"path/filepath"
)

func writeTSConfig(dir string) error {
	src := `{
  "compilerOptions": {
    "target": "ES2022",
    "useDefineForClassFields": true,
    "lib": ["ES2022", "DOM", "DOM.Iterable"],
    "module": "ESNext",
    "skipLibCheck": true,
    "moduleResolution": "bundler",
    "allowImportingTsExtensions": true,
    "resolveJsonModule": true,
    "isolatedModules": true,
    "moduleDetection": "force",
    "noEmit": true,
    "jsx": "react-jsx",
    "strict": true,
    "noUnusedLocals": false,
    "noUnusedParameters": false,
    "noFallthroughCasesInSwitch": true,
    "paths": {
      "@/*": ["./src/*"]
    },
    "baseUrl": "."
  },
  "include": ["src"]
}
`
	return os.WriteFile(filepath.Join(dir, "tsconfig.json"), []byte(src), 0644)
}
