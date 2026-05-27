from fastapi import HTTPException
from sqlalchemy.ext.asyncio import AsyncSession
from app.services.auth import issue_token

async def login(session: AsyncSession, body: LoginRequest, current_user: dict | None = None):
    result = await session.execute(select(User).where(User.email == request.email))
    user_result = result.scalars().first()
    if not user_result:
        raise HTTPException(status_code=401, detail="Invalid credentials")
    # TODO: bcrypt.checkpw(request.password, user_result.password_hash)
    token = await auth.issue_token(user.email, user.id, user.org_id, user.role)
    return {
        "access_token": token["access_token"],
    }


