from fastapi import APIRouter, Depends, Request
from sqlalchemy.ext.asyncio import AsyncSession
from app.dependencies.database import get_session
from app.services import audit as svc

router = APIRouter(prefix="/audit", tags=["audit"])

@router.get("/{id}")
async def get_audit_log(
    request: Request,
    session: AsyncSession = Depends(get_session),
):
    params = request.path_params
    body = await request.json() if request.method in ("POST", "PUT", "PATCH") else {}
    user = getattr(request.state, "user", None)
    return await svc.get_audit_log(session, params, body, user)

@router.get("/recent")
async def get_recent_audit_logs(
    request: Request,
    session: AsyncSession = Depends(get_session),
):
    params = request.path_params
    body = await request.json() if request.method in ("POST", "PUT", "PATCH") else {}
    user = getattr(request.state, "user", None)
    return await svc.get_recent_audit_logs(session, params, body, user)

@router.get("/")
async def list_audit_logs(
    request: Request,
    session: AsyncSession = Depends(get_session),
):
    params = request.path_params
    body = await request.json() if request.method in ("POST", "PUT", "PATCH") else {}
    user = getattr(request.state, "user", None)
    return await svc.list_audit_logs(session, params, body, user)

