from fastapi import HTTPException
from sqlalchemy import select, update, delete
from sqlalchemy.ext.asyncio import AsyncSession

async def get_execution_detail(session: AsyncSession, params: dict, body: dict, user: dict | None = None):
    # @auth execution_log.GetExecutionDetail
    # TODO: integrate OPA policy evaluation
    result = await session.execute(select(ExecutionLog).where(ExecutionLog.ID == request["id"]))
    log = result.scalars().first()
    if not log:
        raise HTTPException(status_code=404, detail="Execution log not found")
    result = await session.execute(select(Workflow).where(Workflow.ID == log["WorkflowID"]))
    wf = result.scalars().first()
    if not wf:
        raise HTTPException(status_code=404, detail="Workflow not found")
    result = await session.execute(select(Organization).where(Organization.ID == log["OrgID"]))
    org = result.scalars().first()
    if not org:
        raise HTTPException(status_code=404, detail="Organization not found")
    detail = await dashboard.build_execution_detail(log["CreditsSpent"], "now", "now", "now", org["Name"], log["Status"], "now", wf["Title"])
    return {
        "detail": detail,
    }


from fastapi import HTTPException
from sqlalchemy import select, update, delete
from sqlalchemy.ext.asyncio import AsyncSession

async def get_execution_report(session: AsyncSession, params: dict, body: dict, user: dict | None = None):
    # @auth execution_log.GetExecutionReport
    # TODO: integrate OPA policy evaluation
    result = await session.execute(select(ExecutionLog).where(ExecutionLog.ID == request["id"]))
    log = result.scalars().first()
    if not log:
        raise HTTPException(status_code=404, detail="Execution log not found")
    return {
        "report_key": log.ReportKey,
    }


