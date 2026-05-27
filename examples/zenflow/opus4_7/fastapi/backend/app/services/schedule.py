from fastapi import HTTPException
from sqlalchemy import select
from sqlalchemy.ext.asyncio import AsyncSession
from app.models.models import Workflow
from app.dependencies.authz import authz_check
from app.services.schedule import build_key, parse_cron
from app.services.session import delete, get, set

async def delete_schedule(session: AsyncSession, id: int, current_user: dict | None = None):
    owner_row = await session.execute(
        select(Workflows.org_id).where(Workflows.id == id)
    )
    owner = owner_row.scalar_one_or_none()
    await authz_check(
        current_user,
        action="DeleteSchedule",
        resource="workflow",
        resource_id=id,
        resource_id=str(id),
        owners={"workflows": {"org_id": owner}},
    )
    result = await session.execute(select(Workflow).where(Workflow.id == id))
    wf = result.scalars().first()
    if not wf:
        raise HTTPException(status_code=404, detail="Workflow not found")
    keyResult = await schedule.build_key(id)
    await session.delete(keyResult.key)


async def get_schedule(session: AsyncSession, id: int, current_user: dict | None = None):
    owner_row = await session.execute(
        select(Workflows.org_id).where(Workflows.id == id)
    )
    owner = owner_row.scalar_one_or_none()
    await authz_check(
        current_user,
        action="GetSchedule",
        resource="workflow",
        resource_id=id,
        resource_id=str(id),
        owners={"workflows": {"org_id": owner}},
    )
    result = await session.execute(select(Workflow).where(Workflow.id == id))
    wf = result.scalars().first()
    if not wf:
        raise HTTPException(status_code=404, detail="Workflow not found")
    keyResult = await schedule.build_key(id)
    sessionResult = await session.get(keyResult.key)
    return {
        "cron": sessionResult["value"],
    }


async def set_schedule(session: AsyncSession, id: int, body: SetScheduleRequest, current_user: dict | None = None):
    owner_row = await session.execute(
        select(Workflows.org_id).where(Workflows.id == id)
    )
    owner = owner_row.scalar_one_or_none()
    await authz_check(
        current_user,
        action="SetSchedule",
        resource="workflow",
        resource_id=id,
        resource_id=str(id),
        owners={"workflows": {"org_id": owner}},
    )
    result = await session.execute(select(Workflow).where(Workflow.id == id))
    wf = result.scalars().first()
    if not wf:
        raise HTTPException(status_code=404, detail="Workflow not found")
    cronResult = await schedule.parse_cron(body.cron)
    keyResult = await schedule.build_key(id)
    await session.set(keyResult.key, 86400, body.cron)
    return {
        "schedule": cronResult,
    }


