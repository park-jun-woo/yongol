from fastapi import HTTPException
from sqlalchemy import select
from sqlalchemy.ext.asyncio import AsyncSession
from app.models.models import Organization
from app.dependencies.authz import authz_check

async def get_dashboard(session: AsyncSession, current_user: dict | None = None):
    await authz_check(
        current_user,
        action="GetDashboard",
        resource="organization",
    )
    result = await session.execute(select(Organization).where(Organization.id == current_user["id"]))
    org = result.scalars().first()
    if not org:
        raise HTTPException(status_code=404, detail="Organization not found")
    summary = await summarize(org.credits_balance, org.name, org.plan_type)
    return {
        "summary": summary,
    }


