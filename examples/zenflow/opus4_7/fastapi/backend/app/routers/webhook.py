from fastapi import APIRouter, Depends, Request
from sqlalchemy.ext.asyncio import AsyncSession
from app.dependencies.database import get_session
from app.services import webhook as svc

router = APIRouter(prefix="/webhook", tags=["webhook"])

@router.post("/")
async def create_webhook(
    request: Request,
    session: AsyncSession = Depends(get_session),
):
    params = request.path_params
    body = await request.json() if request.method in ("POST", "PUT", "PATCH") else {}
    user = getattr(request.state, "user", None)
    return await svc.create_webhook(session, params, body, user)

@router.delete("/{id}")
async def delete_webhook(
    request: Request,
    session: AsyncSession = Depends(get_session),
):
    params = request.path_params
    body = await request.json() if request.method in ("POST", "PUT", "PATCH") else {}
    user = getattr(request.state, "user", None)
    return await svc.delete_webhook(session, params, body, user)

@router.get("/")
async def list_webhooks(
    request: Request,
    session: AsyncSession = Depends(get_session),
):
    params = request.path_params
    body = await request.json() if request.method in ("POST", "PUT", "PATCH") else {}
    user = getattr(request.state, "user", None)
    return await svc.list_webhooks(session, params, body, user)

# Subscribe handler for topic: workflow.executed
async def handle_on_workflow_executed(session: AsyncSession, payload: dict):
    return await svc.on_workflow_executed(session, payload)

