from fastapi import APIRouter, Depends
from sqlalchemy.ext.asyncio import AsyncSession
from app.dependencies.database import get_session
from app.dependencies.auth import get_current_user
from app.services import organization as svc
from app.schemas.organization import VerifyOrgAddressRequest

router = APIRouter(prefix="/organization", tags=["organization"])

@router.post("/{id}/verify-address")
async def verify_org_address(
    id: int,
    body: VerifyOrgAddressRequest,
    current_user: dict = Depends(get_current_user),
    session: AsyncSession = Depends(get_session),
):
    return await svc.verify_org_address(session, id, body, current_user)

