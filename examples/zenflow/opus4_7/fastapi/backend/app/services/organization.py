from fastapi import HTTPException
from sqlalchemy import select, update
from sqlalchemy.ext.asyncio import AsyncSession
from app.models.models import Organization
from app.dependencies.authz import authz_check
from app.services.geocoding import geocode

async def verify_org_address(session: AsyncSession, id: int, body: VerifyOrgAddressRequest, current_user: dict | None = None):
    async with session.begin():
        owner_row = await session.execute(
            select(Organizations.id).where(Organizations.id == id)
        )
        owner = owner_row.scalar_one_or_none()
        await authz_check(
            current_user,
            action="VerifyOrgAddress",
            resource="organization",
            resource_id=str(id),
            owners={"organizations": {"id": owner}},
        )
        result = await session.execute(select(Organization).where(Organization.id == id))
        org = result.scalars().first()
        if not org:
            raise HTTPException(status_code=404, detail="Organization not found")
        geo = await geocode(body.address)
        await session.execute(
            update(Organization).where(Organization.id == org.id).values(address_verified=geo.address_verified, latitude=geo.latitude, longitude=geo.longitude)
        )
        result = await session.execute(select(Organization).where(Organization.id == org.id))
        updated_org = result.scalars().first()
        if not updated_org:
            raise HTTPException(status_code=404, detail="Organization not found")
        return {
            "organization": updated_org,
        }


