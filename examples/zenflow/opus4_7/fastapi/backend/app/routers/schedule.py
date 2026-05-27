from fastapi import APIRouter, Depends
from sqlalchemy.ext.asyncio import AsyncSession
from app.dependencies.database import get_session
from app.dependencies.auth import get_current_user
from app.services import schedule as svc
from app.schemas.schedule import SetScheduleRequest

router = APIRouter(prefix="/schedule", tags=["schedule"])

@router.delete("/{id}/schedule")
async def delete_schedule(
    id: int,
    current_user: dict = Depends(get_current_user),
    session: AsyncSession = Depends(get_session),
):
    return await svc.delete_schedule(session, id, current_user)

@router.get("/{id}/schedule")
async def get_schedule(
    id: int,
    current_user: dict = Depends(get_current_user),
    session: AsyncSession = Depends(get_session),
):
    return await svc.get_schedule(session, id, current_user)

@router.post("/{id}/schedule")
async def set_schedule(
    id: int,
    body: SetScheduleRequest,
    current_user: dict = Depends(get_current_user),
    session: AsyncSession = Depends(get_session),
):
    return await svc.set_schedule(session, id, body, current_user)

