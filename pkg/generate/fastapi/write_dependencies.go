//ff:func feature=gen-fastapi type=util control=sequence
//ff:what writeDependencies — FastAPI 의존성 주입 모듈 파일 기록

package fastapi

import (
	"os"
	"path/filepath"
)

// writeDependencies writes the FastAPI dependency injection modules.
func writeDependencies(appDir string) error {
	depsDir := filepath.Join(appDir, "dependencies")
	if err := os.MkdirAll(depsDir, 0o755); err != nil {
		return err
	}

	// __init__.py
	if err := os.WriteFile(filepath.Join(depsDir, "__init__.py"), []byte(""), 0o644); err != nil {
		return err
	}

	// database.py — session dependency
	dbDep := `from collections.abc import AsyncGenerator

from sqlalchemy.ext.asyncio import AsyncSession

from app.database import async_session


async def get_session() -> AsyncGenerator[AsyncSession, None]:
    async with async_session() as session:
        yield session
`
	if err := os.WriteFile(filepath.Join(depsDir, "database.py"), []byte(dbDep), 0o644); err != nil {
		return err
	}

	// auth.py — current user dependency stub
	authDep := `from fastapi import Depends, HTTPException, status
from fastapi.security import HTTPAuthorizationCredentials, HTTPBearer

security = HTTPBearer()


async def get_current_user(
    credentials: HTTPAuthorizationCredentials = Depends(security),
) -> dict:
    """Validate the bearer token and return the current user dict.

    TODO: Replace with actual JWT / session verification logic.
    """
    token = credentials.credentials
    if not token:
        raise HTTPException(
            status_code=status.HTTP_401_UNAUTHORIZED,
            detail="Not authenticated",
        )
    # TODO: decode and verify token, return user payload
    return {"sub": token}
`
	return os.WriteFile(filepath.Join(depsDir, "auth.py"), []byte(authDep), 0o644)
}
