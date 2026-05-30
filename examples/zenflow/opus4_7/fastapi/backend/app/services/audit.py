from fastapi import HTTPException
from sqlalchemy import select, update, delete
from sqlalchemy.ext.asyncio import AsyncSession

async def get_audit_log(session: AsyncSession, params: dict, body: dict, user: dict | None = None):
    # @auth audit_log.GetAuditLog
    # TODO: integrate OPA policy evaluation
    result = await session.execute(select(AuditLog).where(AuditLog.ID == request["id"]))
    audit_log = result.scalars().first()
    if not audit_log:
        raise HTTPException(status_code=404, detail="Audit log not found")
    return {
        "audit_log": audit_log,
    }


from fastapi import HTTPException
from sqlalchemy import select, update, delete
from sqlalchemy.ext.asyncio import AsyncSession

async def get_recent_audit_logs(session: AsyncSession, params: dict, body: dict, user: dict | None = None):
    # @auth audit_log.GetRecentAuditLogs
    # TODO: integrate OPA policy evaluation
    result = await session.execute(select(AuditLog).where(AuditLog.FilterAction == params, AuditLog.OrgID == currentUser["OrgID"], AuditLog.PageOffset == 0, AuditLog.PerPage == 10))
    items = result.scalars().all()
    return {
        "items": items,
    }


from fastapi import HTTPException
from sqlalchemy import select, update, delete
from sqlalchemy.ext.asyncio import AsyncSession

async def list_audit_logs(session: AsyncSession, params: dict, body: dict, user: dict | None = None):
    # @auth audit_log.ListAuditLogs
    # TODO: integrate OPA policy evaluation
    result = await session.execute(select(AuditLog).where(AuditLog.FilterAction == request["action"], AuditLog.OrgID == currentUser["OrgID"], AuditLog.PageOffset == request["page"], AuditLog.PerPage == request["per_page"]))
    items = result.scalars().all()
    result = await session.execute(select(AuditLog).where(AuditLog.FilterAction == request["action"], AuditLog.OrgID == currentUser["OrgID"]))
    total = result.scalars().first()
    return {
        "items": items,
        "total": total,
    }


