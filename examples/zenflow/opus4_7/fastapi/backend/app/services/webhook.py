from fastapi import HTTPException
from sqlalchemy import select, update, delete
from sqlalchemy.ext.asyncio import AsyncSession

async def create_webhook(session: AsyncSession, params: dict, body: dict, user: dict | None = None):
    async with session.begin():
        # @auth webhook.CreateWebhook
        # TODO: integrate OPA policy evaluation
        webhook = Webhook(EventType=request["event_type"], OrgID=currentUser["OrgID"], Url=request["url"])
        session.add(webhook)
        await session.flush()
        return {
            "webhook": webhook,
        }


from fastapi import HTTPException
from sqlalchemy import select, update, delete
from sqlalchemy.ext.asyncio import AsyncSession

async def delete_webhook(session: AsyncSession, params: dict, body: dict, user: dict | None = None):
    async with session.begin():
        # @auth webhook.DeleteWebhook
        # TODO: integrate OPA policy evaluation
        result = await session.execute(select(Webhook).where(Webhook.ID == request["id"]))
        webhook = result.scalars().first()
        if not webhook:
            raise HTTPException(status_code=404, detail="Webhook not found")
        await session.execute(delete(Webhook).where(Webhook.ID == webhook["ID"]))


from fastapi import HTTPException
from sqlalchemy import select, update, delete
from sqlalchemy.ext.asyncio import AsyncSession

async def list_webhooks(session: AsyncSession, params: dict, body: dict, user: dict | None = None):
    # @auth webhook.ListWebhooks
    # TODO: integrate OPA policy evaluation
    result = await session.execute(select(Webhook).where(Webhook.OrgID == currentUser["OrgID"]))
    webhooks = result.scalars().all()
    return {
        "webhooks": webhooks,
    }


from fastapi import HTTPException
from sqlalchemy import select, update, delete
from sqlalchemy.ext.asyncio import AsyncSession

async def on_workflow_executed(session: AsyncSession, payload: dict):
    await webhookdelivery.deliver(message["Status"], "simulated")


