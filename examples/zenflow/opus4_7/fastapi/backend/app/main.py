from contextlib import asynccontextmanager
from fastapi import FastAPI
from app.database import engine, Base
from app.routers import organization as organization_router
from app.routers import workflow as workflow_router
from app.routers import audit as audit_router
from app.routers import dashboard as dashboard_router
from app.routers import execution as execution_router
from app.routers import schedule as schedule_router
from app.routers import template as template_router
from app.routers import webhook as webhook_router
from app.routers import auth as auth_router


@asynccontextmanager
async def lifespan(app: FastAPI):
    async with engine.begin() as conn:
        await conn.run_sync(Base.metadata.create_all)
    yield
    await engine.dispose()

app = FastAPI(title="zenflow", lifespan=lifespan)

app.include_router(organization_router.router)
app.include_router(workflow_router.router)
app.include_router(audit_router.router)
app.include_router(dashboard_router.router)
app.include_router(execution_router.router)
app.include_router(schedule_router.router)
app.include_router(template_router.router)
app.include_router(webhook_router.router)
app.include_router(auth_router.router)


@app.get("/health")
async def health():
    return {"status": "ok"}
