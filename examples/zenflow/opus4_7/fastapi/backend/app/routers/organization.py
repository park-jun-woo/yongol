from fastapi import APIRouter, Depends, Request
from sqlalchemy.ext.asyncio import AsyncSession
from app.dependencies.database import get_session
from app.services import organization as svc

router = APIRouter(prefix="/organization", tags=["organization"])

@router.post("/{id}/verify-address")
async def verify_org_address(
    request: Request,
    session: AsyncSession = Depends(get_session),
):
    params = request.path_params
    body = await request.json() if request.method in ("POST", "PUT", "PATCH") else {}
    user = getattr(request.state, "user", None)
    return await svc.verify_org_address(session, params, body, user)

