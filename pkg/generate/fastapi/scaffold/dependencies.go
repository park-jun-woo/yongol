//ff:func feature=gen-fastapi type=util control=sequence
//ff:what runtimeDependencies — FastAPI 프로젝트 런타임 의존성 목록

package scaffold

// runtimeDependencies returns the runtime dependencies for a FastAPI project.
func runtimeDependencies() []string {
	return []string{
		"fastapi>=0.110.0",
		"uvicorn[standard]>=0.27.0",
		"sqlalchemy[asyncio]>=2.0.0",
		"asyncpg>=0.29.0",
		"pydantic-settings>=2.0.0",
		"python-jose[cryptography]>=3.3.0",
		"passlib[bcrypt]>=1.7.4",
	}
}
