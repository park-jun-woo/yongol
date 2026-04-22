//ff:func feature=gen-react type=generator control=sequence
//ff:what vite.config.ts — React plugin + @/ alias + /api proxy

package react

import (
	"os"
	"path/filepath"
)

func writeViteConfig(dir string) error {
	src := `import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import path from 'node:path'

export default defineConfig({
  plugins: [react()],
  resolve: {
    alias: {
      '@': path.resolve(__dirname, 'src'),
    },
  },
  server: {
    proxy: {
      '/api': 'http://localhost:8080',
    },
  },
})
`
	return os.WriteFile(filepath.Join(dir, "vite.config.ts"), []byte(src), 0644)
}
