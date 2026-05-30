from fastapi import APIRouter, Depends, Request
from sqlalchemy.ext.asyncio import AsyncSession
from app.dependencies.database import get_session
from app.services import dashboard as svc

router = APIRouter(prefix="/dashboard", tags=["dashboard"])

@router.get("/")
async def get_dashboard(
    request: Request,
    session: AsyncSession = Depends(get_session),
):
    params = request.path_params
    body = await request.json() if request.method in ("POST", "PUT", "PATCH") else {}
    user = getattr(request.state, "user", None)
    return await svc.get_dashboard(session, params, body, user)

