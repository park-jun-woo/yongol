//ff:func feature=gen-fastapi type=util control=sequence
//ff:what devDependencies — FastAPI 프로젝트 개발 의존성 목록

package scaffold

// devDependencies returns the dev dependencies for a FastAPI project.
func devDependencies() []string {
	return []string{
		"ruff>=0.3.0",
		"mypy>=1.8.0",
		"pytest>=8.0.0",
		"pytest-asyncio>=0.23.0",
		"types-python-jose>=3.3.0",
	}
}
