from fastapi import APIRouter, Depends
from sqlalchemy.ext.asyncio import AsyncSession
from app.dependencies.database import get_session
from app.dependencies.auth import get_current_user
from app.services import webhook as svc
from app.schemas.webhook import CreateWebhookRequest

router = APIRouter(prefix="/webhook", tags=["webhook"])

@router.post("/")
async def create_webhook(
    body: CreateWebhookRequest,
    current_user: dict = Depends(get_current_user),
    session: AsyncSession = Depends(get_session),
):
    return await svc.create_webhook(session, body, current_user)

@router.delete("/{id}")
async def delete_webhook(
    id: int,
    current_user: dict = Depends(get_current_user),
    session: AsyncSession = Depends(get_session),
):
    return await svc.delete_webhook(session, id, current_user)

@router.get("/")
async def list_webhooks(
    current_user: dict = Depends(get_current_user),
    session: AsyncSession = Depends(get_session),
):
    return await svc.list_webhooks(session, current_user)

# Subscribe handler for topic: workflow.executed
async def handle_on_workflow_executed(session: AsyncSession, payload: dict):
    return await svc.on_workflow_executed(session, payload)

