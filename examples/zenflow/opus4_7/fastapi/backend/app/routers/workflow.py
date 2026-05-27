from fastapi import APIRouter, Depends
from sqlalchemy.ext.asyncio import AsyncSession
from app.dependencies.database import get_session
from app.dependencies.auth import get_current_user
from app.services import workflow as svc
from app.schemas.workflow import AddActionRequest, CreateWorkflowRequest, SaveWorkflowActionsRequest

router = APIRouter(prefix="/workflow", tags=["workflow"])

@router.post("/{id}/activate")
async def activate_workflow(
    id: int,
    current_user: dict = Depends(get_current_user),
    session: AsyncSession = Depends(get_session),
):
    return await svc.activate_workflow(session, id, current_user)

@router.post("/{id}/actions")
async def add_action(
    id: int,
    body: AddActionRequest,
    current_user: dict = Depends(get_current_user),
    session: AsyncSession = Depends(get_session),
):
    return await svc.add_action(session, id, body, current_user)

@router.post("/{id}/archive")
async def archive_workflow(
    id: int,
    current_user: dict = Depends(get_current_user),
    session: AsyncSession = Depends(get_session),
):
    return await svc.archive_workflow(session, id, current_user)

@router.post("/{id}/auto-assign")
async def auto_assign_workflow(
    id: int,
    current_user: dict = Depends(get_current_user),
    session: AsyncSession = Depends(get_session),
):
    return await svc.auto_assign_workflow(session, id, current_user)

@router.post("/")
async def create_workflow(
    body: CreateWorkflowRequest,
    current_user: dict = Depends(get_current_user),
    session: AsyncSession = Depends(get_session),
):
    return await svc.create_workflow(session, body, current_user)

@router.post("/{id}/new-version")
async def create_workflow_version(
    id: int,
    current_user: dict = Depends(get_current_user),
    session: AsyncSession = Depends(get_session),
):
    return await svc.create_workflow_version(session, id, current_user)

@router.post("/{id}/execute-with-report")
async def execute_with_report(
    id: int,
    current_user: dict = Depends(get_current_user),
    session: AsyncSession = Depends(get_session),
):
    return await svc.execute_with_report(session, id, current_user)

@router.post("/{id}/execute")
async def execute_workflow(
    id: int,
    current_user: dict = Depends(get_current_user),
    session: AsyncSession = Depends(get_session),
):
    return await svc.execute_workflow(session, id, current_user)

@router.get("/{id}")
async def get_workflow(
    id: int,
    current_user: dict = Depends(get_current_user),
    session: AsyncSession = Depends(get_session),
):
    return await svc.get_workflow(session, id, current_user)

@router.get("/{id}/execution-logs")
async def list_execution_logs(
    id: int,
    current_user: dict = Depends(get_current_user),
    session: AsyncSession = Depends(get_session),
):
    return await svc.list_execution_logs(session, id, current_user)

@router.get("/{id}/versions")
async def list_workflow_versions(
    id: int,
    current_user: dict = Depends(get_current_user),
    session: AsyncSession = Depends(get_session),
):
    return await svc.list_workflow_versions(session, id, current_user)

@router.get("/")
async def list_workflows(
    current_user: dict = Depends(get_current_user),
    session: AsyncSession = Depends(get_session),
):
    return await svc.list_workflows(session, current_user)

@router.post("/{id}/pause")
async def pause_workflow(
    id: int,
    current_user: dict = Depends(get_current_user),
    session: AsyncSession = Depends(get_session),
):
    return await svc.pause_workflow(session, id, current_user)

@router.put("/{id}/actions")
async def save_workflow_actions(
    id: int,
    body: SaveWorkflowActionsRequest,
    current_user: dict = Depends(get_current_user),
    session: AsyncSession = Depends(get_session),
):
    return await svc.save_workflow_actions(session, id, body, current_user)

