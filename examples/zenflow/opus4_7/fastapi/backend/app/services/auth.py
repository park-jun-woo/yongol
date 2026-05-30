from fastapi import HTTPException
from sqlalchemy import select, update, delete
from sqlalchemy.ext.asyncio import AsyncSession

async def login(session: AsyncSession, params: dict, body: dict, user: dict | None = None):
    result = await session.execute(select(User).where(User.email == request.email))
    user = result.scalars().first()
    if not user:
        raise HTTPException(status_code=401, detail="Invalid credentials")
    # TODO: bcrypt.checkpw(request.password, user.password_hash)
    token = await auth.issue_token(user["Email"], user["ID"], user["OrgID"], user["Role"])
    return {
        "access_token": token.AccessToken,
    }


