from fastapi import HTTPException
from sqlalchemy import select, update, delete
from sqlalchemy.ext.asyncio import AsyncSession

async def get_dashboard(session: AsyncSession, params: dict, body: dict, user: dict | None = None):
    # @auth organization.GetDashboard
    # TODO: integrate OPA policy evaluation
    result = await session.execute(select(Organization).where(Organization.ID == currentUser["OrgID"]))
    org = result.scalars().first()
    if not org:
        raise HTTPException(status_code=404, detail="Organization not found")
    summary = await dashboard.summarize(org["CreditsBalance"], org["Name"], org["PlanType"])
    return {
        "summary": summary,
    }


