from fastapi import APIRouter, Depends
from sqlalchemy.ext.asyncio import AsyncSession
from app.dependencies.database import get_session
from app.dependencies.auth import get_current_user
from app.services import execution as svc

router = APIRouter(prefix="/execution", tags=["execution"])

@router.get("/{id}/detail")
async def get_execution_detail(
    id: int,
    current_user: dict = Depends(get_current_user),
    session: AsyncSession = Depends(get_session),
):
    return await svc.get_execution_detail(session, id, current_user)

@router.get("/{id}/report")
async def get_execution_report(
    id: int,
    current_user: dict = Depends(get_current_user),
    session: AsyncSession = Depends(get_session),
):
    return await svc.get_execution_report(session, id, current_user)

