//ff:func feature=gen-fastapi type=generator control=sequence
//ff:what RenderEnvExample — FastAPI .env.example 생성

package scaffold

import "strings"

// RenderEnvExample produces the .env.example content with default settings.
func RenderEnvExample() (string, error) {
	var b strings.Builder

	b.WriteString("DATABASE_URL=postgresql+asyncpg://postgres:postgres@localhost:5432/app\n")
	b.WriteString("JWT_SECRET=change-me-in-production\n")
	b.WriteString("JWT_ALGORITHM=HS256\n")
	b.WriteString("JWT_EXPIRE_MINUTES=30\n")
	b.WriteString("DEBUG=false\n")

	return b.String(), nil
}
