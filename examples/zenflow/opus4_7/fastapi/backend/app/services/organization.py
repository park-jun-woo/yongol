from fastapi import HTTPException
from sqlalchemy import select, update, delete
from sqlalchemy.ext.asyncio import AsyncSession

async def verify_org_address(session: AsyncSession, params: dict, body: dict, user: dict | None = None):
    async with session.begin():
        # @auth organization.VerifyOrgAddress
        # TODO: integrate OPA policy evaluation
        result = await session.execute(select(Organization).where(Organization.ID == request["id"]))
        org = result.scalars().first()
        if not org:
            raise HTTPException(status_code=404, detail="Organization not found")
        geo = await geocoding.geocode(request["address"])
        await session.execute(
            update(Organization).where(Organization.id == params["id"]).values(**body)
        )
        result = await session.execute(select(Organization).where(Organization.ID == org["ID"]))
        updated_org = result.scalars().first()
        if not updated_org:
            raise HTTPException(status_code=404, detail="Organization not found")
        return {
            "organization": updated_org,
        }


