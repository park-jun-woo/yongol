from fastapi import APIRouter, Depends
from sqlalchemy.ext.asyncio import AsyncSession
from app.dependencies.database import get_session
from app.dependencies.auth import get_current_user
from app.services import dashboard as svc

router = APIRouter(prefix="/dashboard", tags=["dashboard"])

@router.get("/")
async def get_dashboard(
    current_user: dict = Depends(get_current_user),
    session: AsyncSession = Depends(get_session),
):
    return await svc.get_dashboard(session, current_user)

