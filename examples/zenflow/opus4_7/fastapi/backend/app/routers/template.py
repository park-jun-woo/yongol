from fastapi import APIRouter, Depends, Request
from sqlalchemy.ext.asyncio import AsyncSession
from app.dependencies.database import get_session
from app.services import template as svc

router = APIRouter(prefix="/template", tags=["template"])

@router.post("/{id}/clone")
async def clone_template(
    request: Request,
    session: AsyncSession = Depends(get_session),
):
    params = request.path_params
    body = await request.json() if request.method in ("POST", "PUT", "PATCH") else {}
    user = getattr(request.state, "user", None)
    return await svc.clone_template(session, params, body, user)

@router.get("/{id}")
async def get_template(
    request: Request,
    session: AsyncSession = Depends(get_session),
):
    params = request.path_params
    body = await request.json() if request.method in ("POST", "PUT", "PATCH") else {}
    user = getattr(request.state, "user", None)
    return await svc.get_template(session, params, body, user)

@router.get("/")
async def list_templates(
    request: Request,
    session: AsyncSession = Depends(get_session),
):
    params = request.path_params
    body = await request.json() if request.method in ("POST", "PUT", "PATCH") else {}
    user = getattr(request.state, "user", None)
    return await svc.list_templates(session, params, body, user)

@router.post("/")
async def publish_template(
    request: Request,
    session: AsyncSession = Depends(get_session),
):
    params = request.path_params
    body = await request.json() if request.method in ("POST", "PUT", "PATCH") else {}
    user = getattr(request.state, "user", None)
    return await svc.publish_template(session, params, body, user)

