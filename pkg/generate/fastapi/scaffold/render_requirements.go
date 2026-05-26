//ff:func feature=gen-fastapi type=generator control=iteration dimension=1
//ff:what RenderRequirements — FastAPI requirements.txt 생성

package scaffold

import (
	"strings"
)

// RenderRequirements produces the requirements.txt content for a FastAPI project.
func RenderRequirements() (string, error) {
	var b strings.Builder

	for _, dep := range runtimeDependencies() {
		b.WriteString(dep + "\n")
	}

	return b.String(), nil
}
