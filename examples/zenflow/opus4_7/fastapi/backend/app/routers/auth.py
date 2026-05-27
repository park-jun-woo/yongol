from fastapi import APIRouter, Depends
from sqlalchemy.ext.asyncio import AsyncSession
from app.dependencies.database import get_session
from app.dependencies.auth import get_current_user
from app.services import auth as svc
from app.schemas.auth import LoginRequest

router = APIRouter(prefix="/auth", tags=["auth"])

@router.post("/login")
async def login(
    body: LoginRequest,
    current_user: dict = Depends(get_current_user),
    session: AsyncSession = Depends(get_session),
):
    return await svc.login(session, body, current_user)

