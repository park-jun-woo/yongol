//ff:func feature=gen-fastapi type=generator control=sequence
//ff:what RenderConfig — FastAPI config.py (pydantic-settings) Python 소스 생성

package boot

import (
	"strings"
)

// RenderConfig produces the config.py content using pydantic-settings
// BaseSettings for environment variable management.
func RenderConfig() (string, error) {
	var b strings.Builder

	b.WriteString("from pydantic_settings import BaseSettings\n\n\n")
	b.WriteString("class Settings(BaseSettings):\n")
	b.WriteString("    database_url: str = \"postgresql+asyncpg://localhost/app\"\n")
	b.WriteString("    jwt_secret: str = \"change-me\"\n")
	b.WriteString("    jwt_algorithm: str = \"HS256\"\n")
	b.WriteString("    jwt_expire_minutes: int = 30\n")
	b.WriteString("    debug: bool = False\n\n")
	b.WriteString("    model_config = {\"env_file\": \".env\"}\n\n\n")
	b.WriteString("settings = Settings()\n")

	return b.String(), nil
}
