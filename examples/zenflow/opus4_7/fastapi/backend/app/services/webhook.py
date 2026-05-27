from fastapi import HTTPException
from sqlalchemy import select, delete
from sqlalchemy.ext.asyncio import AsyncSession
from app.models.models import Webhook
from app.dependencies.authz import authz_check
from app.services.webhookdelivery import deliver

async def create_webhook(session: AsyncSession, body: CreateWebhookRequest, current_user: dict | None = None):
    async with session.begin():
        await authz_check(
            current_user,
            action="CreateWebhook",
            resource="webhook",
        )
        webhook = Webhook(event_type=body.event_type, org_id=current_user["org_id"], url=body.url)
        session.add(webhook)
        await session.flush()
        return {
            "webhook": webhook,
        }


async def delete_webhook(session: AsyncSession, id: int, current_user: dict | None = None):
    async with session.begin():
        owner_row = await session.execute(
            select(Webhooks.org_id).where(Webhooks.id == id)
        )
        owner = owner_row.scalar_one_or_none()
        await authz_check(
            current_user,
            action="DeleteWebhook",
            resource="webhook",
            resource_id=str(id),
            owners={"webhooks": {"org_id": owner}},
        )
        result = await session.execute(select(Webhook).where(Webhook.id == id))
        webhook = result.scalars().first()
        if not webhook:
            raise HTTPException(status_code=404, detail="Webhook not found")
        await session.execute(delete(Webhook).where(Webhook.id == webhook.id))


async def list_webhooks(session: AsyncSession, current_user: dict | None = None):
    await authz_check(
        current_user,
        action="ListWebhooks",
        resource="webhook",
    )
    result = await session.execute(select(Webhook).where(Webhook.org_id == current_user["org_id"]))
    webhooks = result.scalars().all()
    return {
        "webhooks": webhooks,
    }


async def on_workflow_executed(session: AsyncSession, payload: dict):
    message = payload
    await deliver(message.status, "simulated")


