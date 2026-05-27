from fastapi import HTTPException
from sqlalchemy import select
from sqlalchemy.ext.asyncio import AsyncSession
from app.models.models import AuditLog
from app.dependencies.authz import authz_check

async def get_audit_log(session: AsyncSession, id: int, current_user: dict | None = None):
    owner_row = await session.execute(
        select(AuditLogs.org_id).where(AuditLogs.id == id)
    )
    owner = owner_row.scalar_one_or_none()
    await authz_check(
        current_user,
        action="GetAuditLog",
        resource="audit_log",
        resource_id=str(id),
        owners={"audit_logs": {"org_id": owner}},
    )
    result = await session.execute(select(AuditLog).where(AuditLog.id == id))
    audit_log = result.scalars().first()
    if not audit_log:
        raise HTTPException(status_code=404, detail="Audit log not found")
    return {
        "audit_log": audit_log,
    }


async def get_recent_audit_logs(session: AsyncSession, current_user: dict | None = None):
    await authz_check(
        current_user,
        action="GetRecentAuditLogs",
        resource="audit_log",
    )
    result = await session.execute(select(AuditLog).where(AuditLog.filter_action == params, AuditLog.org_id == current_user["org_id"]).offset(0).limit(10))
    items = result.scalars().all()
    return {
        "items": items,
    }


async def list_audit_logs(session: AsyncSession, action: str | None = None, page: int | None = None, per_page: int | None = None, current_user: dict | None = None):
    await authz_check(
        current_user,
        action="ListAuditLogs",
        resource="audit_log",
    )
    result = await session.execute(select(AuditLog).where(AuditLog.filter_action == action, AuditLog.org_id == current_user["org_id"]).offset(page).limit(per_page))
    items = result.scalars().all()
    result = await session.execute(select(AuditLog).where(AuditLog.filter_action == action, AuditLog.org_id == current_user["org_id"]))
    total = result.scalars().first()
    return {
        "items": items,
        "total": total,
    }


