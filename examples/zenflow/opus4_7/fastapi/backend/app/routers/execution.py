from fastapi import APIRouter, Depends, Request
from sqlalchemy.ext.asyncio import AsyncSession
from app.dependencies.database import get_session
from app.services import execution as svc

router = APIRouter(prefix="/execution", tags=["execution"])

@router.get("/{id}/detail")
async def get_execution_detail(
    request: Request,
    session: AsyncSession = Depends(get_session),
):
    params = request.path_params
    body = await request.json() if request.method in ("POST", "PUT", "PATCH") else {}
    user = getattr(request.state, "user", None)
    return await svc.get_execution_detail(session, params, body, user)

@router.get("/{id}/report")
async def get_execution_report(
    request: Request,
    session: AsyncSession = Depends(get_session),
):
    params = request.path_params
    body = await request.json() if request.method in ("POST", "PUT", "PATCH") else {}
    user = getattr(request.state, "user", None)
    return await svc.get_execution_report(session, params, body, user)

