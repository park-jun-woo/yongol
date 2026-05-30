from fastapi import HTTPException
from sqlalchemy import select, update, delete
from sqlalchemy.ext.asyncio import AsyncSession

async def delete_schedule(session: AsyncSession, params: dict, body: dict, user: dict | None = None):
    # @auth workflow.DeleteSchedule
    # TODO: integrate OPA policy evaluation
    result = await session.execute(select(Workflow).where(Workflow.ID == request["id"]))
    wf = result.scalars().first()
    if not wf:
        raise HTTPException(status_code=404, detail="Workflow not found")
    keyResult = await schedule.build_key(request["id"])
    await session.delete(keyResult["Key"])


from fastapi import HTTPException
from sqlalchemy import select, update, delete
from sqlalchemy.ext.asyncio import AsyncSession

async def get_schedule(session: AsyncSession, params: dict, body: dict, user: dict | None = None):
    # @auth workflow.GetSchedule
    # TODO: integrate OPA policy evaluation
    result = await session.execute(select(Workflow).where(Workflow.ID == request["id"]))
    wf = result.scalars().first()
    if not wf:
        raise HTTPException(status_code=404, detail="Workflow not found")
    keyResult = await schedule.build_key(request["id"])
    sessionResult = await session.get(keyResult["Key"])
    return {
        "cron": sessionResult.Value,
    }


from fastapi import HTTPException
from sqlalchemy import select, update, delete
from sqlalchemy.ext.asyncio import AsyncSession

async def set_schedule(session: AsyncSession, params: dict, body: dict, user: dict | None = None):
    # @auth workflow.SetSchedule
    # TODO: integrate OPA policy evaluation
    result = await session.execute(select(Workflow).where(Workflow.ID == request["id"]))
    wf = result.scalars().first()
    if not wf:
        raise HTTPException(status_code=404, detail="Workflow not found")
    cronResult = await schedule.parse_cron(request["cron"])
    keyResult = await schedule.build_key(request["id"])
    await session.set(keyResult["Key"], 86400, request["cron"])
    return {
        "schedule": cronResult,
    }


