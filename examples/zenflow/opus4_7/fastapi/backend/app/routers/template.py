from fastapi import APIRouter, Depends
from sqlalchemy.ext.asyncio import AsyncSession
from app.dependencies.database import get_session
from app.dependencies.auth import get_current_user
from app.services import template as svc

router = APIRouter(prefix="/template", tags=["template"])

@router.post("/{id}/clone")
async def clone_template(
    id: int,
    current_user: dict = Depends(get_current_user),
    session: AsyncSession = Depends(get_session),
):
    return await svc.clone_template(session, id, current_user)

@router.get("/{id}")
async def get_template(
    id: int,
    current_user: dict = Depends(get_current_user),
    session: AsyncSession = Depends(get_session),
):
    return await svc.get_template(session, id, current_user)

@router.get("/")
async def list_templates(
    category: str | None = None,
    cursor: str | None = None,
    per_page: int | None = None,
    current_user: dict = Depends(get_current_user),
    session: AsyncSession = Depends(get_session),
):
    return await svc.list_templates(session, category, cursor, per_page, current_user)

@router.post("/")
async def publish_template(
    body: PublishTemplateRequest,
    current_user: dict = Depends(get_current_user),
    session: AsyncSession = Depends(get_session),
):
    return await svc.publish_template(session, body, current_user)

