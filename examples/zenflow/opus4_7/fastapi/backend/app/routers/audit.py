from fastapi import APIRouter, Depends
from sqlalchemy.ext.asyncio import AsyncSession
from app.dependencies.database import get_session
from app.dependencies.auth import get_current_user
from app.services import audit as svc

router = APIRouter(prefix="/audit", tags=["audit"])

@router.get("/{id}")
async def get_audit_log(
    id: int,
    current_user: dict = Depends(get_current_user),
    session: AsyncSession = Depends(get_session),
):
    return await svc.get_audit_log(session, id, current_user)

@router.get("/recent")
async def get_recent_audit_logs(
    current_user: dict = Depends(get_current_user),
    session: AsyncSession = Depends(get_session),
):
    return await svc.get_recent_audit_logs(session, current_user)

@router.get("/")
async def list_audit_logs(
    action: str | None = None,
    page: int | None = None,
    per_page: int | None = None,
    current_user: dict = Depends(get_current_user),
    session: AsyncSession = Depends(get_session),
):
    return await svc.list_audit_logs(session, action, page, per_page, current_user)

