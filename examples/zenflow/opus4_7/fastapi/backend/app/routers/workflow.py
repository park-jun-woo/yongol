from fastapi import APIRouter, Depends, Request
from sqlalchemy.ext.asyncio import AsyncSession
from app.dependencies.database import get_session
from app.services import workflow as svc

router = APIRouter(prefix="/workflow", tags=["workflow"])

@router.post("/{id}/activate")
async def activate_workflow(
    request: Request,
    session: AsyncSession = Depends(get_session),
):
    params = request.path_params
    body = await request.json() if request.method in ("POST", "PUT", "PATCH") else {}
    user = getattr(request.state, "user", None)
    return await svc.activate_workflow(session, params, body, user)

@router.post("/{id}/actions")
async def add_action(
    request: Request,
    session: AsyncSession = Depends(get_session),
):
    params = request.path_params
    body = await request.json() if request.method in ("POST", "PUT", "PATCH") else {}
    user = getattr(request.state, "user", None)
    return await svc.add_action(session, params, body, user)

@router.post("/{id}/archive")
async def archive_workflow(
    request: Request,
    session: AsyncSession = Depends(get_session),
):
    params = request.path_params
    body = await request.json() if request.method in ("POST", "PUT", "PATCH") else {}
    user = getattr(request.state, "user", None)
    return await svc.archive_workflow(session, params, body, user)

@router.post("/{id}/auto-assign")
async def auto_assign_workflow(
    request: Request,
    session: AsyncSession = Depends(get_session),
):
    params = request.path_params
    body = await request.json() if request.method in ("POST", "PUT", "PATCH") else {}
    user = getattr(request.state, "user", None)
    return await svc.auto_assign_workflow(session, params, body, user)

@router.post("/")
async def create_workflow(
    request: Request,
    session: AsyncSession = Depends(get_session),
):
    params = request.path_params
    body = await request.json() if request.method in ("POST", "PUT", "PATCH") else {}
    user = getattr(request.state, "user", None)
    return await svc.create_workflow(session, params, body, user)

@router.post("/{id}/new-version")
async def create_workflow_version(
    request: Request,
    session: AsyncSession = Depends(get_session),
):
    params = request.path_params
    body = await request.json() if request.method in ("POST", "PUT", "PATCH") else {}
    user = getattr(request.state, "user", None)
    return await svc.create_workflow_version(session, params, body, user)

@router.post("/{id}/execute-with-report")
async def execute_with_report(
    request: Request,
    session: AsyncSession = Depends(get_session),
):
    params = request.path_params
    body = await request.json() if request.method in ("POST", "PUT", "PATCH") else {}
    user = getattr(request.state, "user", None)
    return await svc.execute_with_report(session, params, body, user)

@router.post("/{id}/execute")
async def execute_workflow(
    request: Request,
    session: AsyncSession = Depends(get_session),
):
    params = request.path_params
    body = await request.json() if request.method in ("POST", "PUT", "PATCH") else {}
    user = getattr(request.state, "user", None)
    return await svc.execute_workflow(session, params, body, user)

@router.get("/{id}")
async def get_workflow(
    request: Request,
    session: AsyncSession = Depends(get_session),
):
    params = request.path_params
    body = await request.json() if request.method in ("POST", "PUT", "PATCH") else {}
    user = getattr(request.state, "user", None)
    return await svc.get_workflow(session, params, body, user)

@router.get("/{id}/execution-logs")
async def list_execution_logs(
    request: Request,
    session: AsyncSession = Depends(get_session),
):
    params = request.path_params
    body = await request.json() if request.method in ("POST", "PUT", "PATCH") else {}
    user = getattr(request.state, "user", None)
    return await svc.list_execution_logs(session, params, body, user)

@router.get("/{id}/versions")
async def list_workflow_versions(
    request: Request,
    session: AsyncSession = Depends(get_session),
):
    params = request.path_params
    body = await request.json() if request.method in ("POST", "PUT", "PATCH") else {}
    user = getattr(request.state, "user", None)
    return await svc.list_workflow_versions(session, params, body, user)

@router.get("/")
async def list_workflows(
    request: Request,
    session: AsyncSession = Depends(get_session),
):
    params = request.path_params
    body = await request.json() if request.method in ("POST", "PUT", "PATCH") else {}
    user = getattr(request.state, "user", None)
    return await svc.list_workflows(session, params, body, user)

@router.post("/{id}/pause")
async def pause_workflow(
    request: Request,
    session: AsyncSession = Depends(get_session),
):
    params = request.path_params
    body = await request.json() if request.method in ("POST", "PUT", "PATCH") else {}
    user = getattr(request.state, "user", None)
    return await svc.pause_workflow(session, params, body, user)

@router.put("/{id}/actions")
async def save_workflow_actions(
    request: Request,
    session: AsyncSession = Depends(get_session),
):
    params = request.path_params
    body = await request.json() if request.method in ("POST", "PUT", "PATCH") else {}
    user = getattr(request.state, "user", None)
    return await svc.save_workflow_actions(session, params, body, user)

