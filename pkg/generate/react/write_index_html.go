//ff:func feature=gen-react type=generator control=sequence
//ff:what index.html 파일을 생성한다

package react

import (
	"os"
	"path/filepath"
)

func writeIndexHTML(dir string) error {
	src := `<!DOCTYPE html>
<html lang="ko" class="h-full">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>App</title>
  </head>
  <body class="min-h-full bg-background text-foreground antialiased">
    <div id="root"></div>
    <script type="module" src="/src/main.tsx"></script>
  </body>
</html>
`
	return os.WriteFile(filepath.Join(dir, "index.html"), []byte(src), 0644)
}
