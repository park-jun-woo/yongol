from fastapi import HTTPException
from sqlalchemy import select
from sqlalchemy.ext.asyncio import AsyncSession
from app.models.models import ExecutionLog, Organization, Workflow
from app.dependencies.authz import authz_check
from app.services.dashboard import build_execution_detail

async def get_execution_detail(session: AsyncSession, id: int, current_user: dict | None = None):
    owner_row = await session.execute(
        select(ExecutionLogs.org_id).where(ExecutionLogs.id == id)
    )
    owner = owner_row.scalar_one_or_none()
    await authz_check(
        current_user,
        action="GetExecutionDetail",
        resource="execution_log",
        resource_id=str(id),
        owners={"execution_logs": {"org_id": owner}},
    )
    result = await session.execute(select(ExecutionLog).where(ExecutionLog.id == id))
    log = result.scalars().first()
    if not log:
        raise HTTPException(status_code=404, detail="Execution log not found")
    result = await session.execute(select(Workflow).where(Workflow.id == log.id))
    wf = result.scalars().first()
    if not wf:
        raise HTTPException(status_code=404, detail="Workflow not found")
    result = await session.execute(select(Organization).where(Organization.id == log.id))
    org = result.scalars().first()
    if not org:
        raise HTTPException(status_code=404, detail="Organization not found")
    detail = await build_execution_detail(log.credits_spent, "now", "now", "now", org.name, log.status, "now", wf.title)
    return {
        "detail": detail,
    }


async def get_execution_report(session: AsyncSession, id: int, current_user: dict | None = None):
    await authz_check(
        current_user,
        action="GetExecutionReport",
        resource="execution_log",
    )
    result = await session.execute(select(ExecutionLog).where(ExecutionLog.id == id))
    log = result.scalars().first()
    if not log:
        raise HTTPException(status_code=404, detail="Execution log not found")
    return {
        "report_key": log["report_key"],
    }


