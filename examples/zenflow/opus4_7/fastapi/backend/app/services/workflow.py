from fastapi import HTTPException
from sqlalchemy import select, update, delete
from sqlalchemy.ext.asyncio import AsyncSession
from app.models.models import Action, ExecutionLog, Organization, User, Workflow
from app.dependencies.event_bus import EventBus
from app.dependencies.authz import authz_check
from app.services.billing import is_zero_balance
from app.services.report import generate_report
from app.services.version import next_version, resolve_root_id
from app.services.worker import process_action

async def activate_workflow(session: AsyncSession, id: int, current_user: dict | None = None):
    async with session.begin():
        owner_row = await session.execute(
            select(Workflows.org_id).where(Workflows.id == id)
        )
        owner = owner_row.scalar_one_or_none()
        await authz_check(
            current_user,
            action="ActivateWorkflow",
            resource="workflow",
            resource_id=str(id),
            owners={"workflows": {"org_id": owner}},
        )
        result = await session.execute(select(Workflow).where(Workflow.id == id))
        wf = result.scalars().first()
        if not wf:
            raise HTTPException(status_code=404, detail="Workflow not found")
        result = await session.execute(select(Organization).where(Organization.id == wf.id))
        org = result.scalars().first()
        if not org:
            raise HTTPException(status_code=404, detail="Organization not found")
        if is_zero_balance(org.credits_balance):
            raise HTTPException(status_code=402, detail="Insufficient credits")
        # @state workflows.ActivateWorkflow — transition guard
        allowed_activate_workflow: dict[str, bool] = {
            "draft": True,
            "paused": True,
        }
        if wf.status not in allowed_activate_workflow:
            raise HTTPException(status_code=409, detail="Cannot activate workflow")
        await session.execute(
            update(Workflow).where(Workflow.id == wf.id).values(status="active")
        )
        result = await session.execute(select(Workflow).where(Workflow.id == wf.id))
        updated = result.scalars().first()
        return {
            "workflow": updated,
        }


async def add_action(session: AsyncSession, id: int, body: AddActionRequest, current_user: dict | None = None):
    async with session.begin():
        owner_row = await session.execute(
            select(Workflows.org_id).where(Workflows.id == id)
        )
        owner = owner_row.scalar_one_or_none()
        await authz_check(
            current_user,
            action="AddAction",
            resource="workflow",
            resource_id=str(id),
            owners={"workflows": {"org_id": owner}},
        )
        result = await session.execute(select(Workflow).where(Workflow.id == id))
        wf = result.scalars().first()
        if not wf:
            raise HTTPException(status_code=404, detail="Workflow not found")
        action = Action(action_type=body.action_type, config=body.config, sequence_order=body.sequence_order, workflow_id=wf.workflow_id)
        session.add(action)
        await session.flush()
        return {
            "action": action,
        }


async def archive_workflow(session: AsyncSession, id: int, current_user: dict | None = None):
    async with session.begin():
        owner_row = await session.execute(
            select(Workflows.org_id).where(Workflows.id == id)
        )
        owner = owner_row.scalar_one_or_none()
        await authz_check(
            current_user,
            action="ArchiveWorkflow",
            resource="workflow",
            resource_id=str(id),
            owners={"workflows": {"org_id": owner}},
        )
        result = await session.execute(select(Workflow).where(Workflow.id == id))
        wf = result.scalars().first()
        if not wf:
            raise HTTPException(status_code=404, detail="Workflow not found")
        # @state workflows.ArchiveWorkflow — transition guard
        allowed_archive_workflow: dict[str, bool] = {
            "active": True,
        }
        if wf.status not in allowed_archive_workflow:
            raise HTTPException(status_code=409, detail="Cannot archive workflow")
        await session.execute(
            update(Workflow).where(Workflow.id == wf.id).values(status="archived")
        )
        result = await session.execute(select(Workflow).where(Workflow.id == wf.id))
        updated = result.scalars().first()
        return {
            "workflow": updated,
        }


async def auto_assign_workflow(session: AsyncSession, id: int, current_user: dict | None = None):
    async with session.begin():
        owner_row = await session.execute(
            select(Workflows.org_id).where(Workflows.id == id)
        )
        owner = owner_row.scalar_one_or_none()
        await authz_check(
            current_user,
            action="AutoAssignWorkflow",
            resource="workflow",
            resource_id=str(id),
            owners={"workflows": {"org_id": owner}},
        )
        result = await session.execute(select(Workflow).where(Workflow.id == id))
        wf = result.scalars().first()
        if not wf:
            raise HTTPException(status_code=404, detail="Workflow not found")
        result = await session.execute(select(User).where(User.org_id == wf.org_id))
        memberCount = result.scalars().first()
        match = await match_member(memberCount, wf.trigger_event)
        await session.execute(
            update(Workflow).where(Workflow.id == wf.id).values(confidence=match.confidence, member_id=match.member_id)
        )
        result = await session.execute(select(Workflow).where(Workflow.id == wf.id))
        updated = result.scalars().first()
        if not updated:
            raise HTTPException(status_code=404, detail="Workflow not found")
        return {
            "workflow": updated,
        }


async def create_workflow(session: AsyncSession, body: CreateWorkflowRequest, current_user: dict | None = None):
    async with session.begin():
        await authz_check(
            current_user,
            action="CreateWorkflow",
            resource="workflow",
        )
        wf = Workflow(org_id=current_user["org_id"], title=body.title, trigger_event=body.trigger_event)
        session.add(wf)
        await session.flush()
        return {
            "workflow": wf,
        }


async def create_workflow_version(session: AsyncSession, id: int, current_user: dict | None = None):
    async with session.begin():
        owner_row = await session.execute(
            select(Workflows.org_id).where(Workflows.id == id)
        )
        owner = owner_row.scalar_one_or_none()
        await authz_check(
            current_user,
            action="CreateWorkflowVersion",
            resource="workflow",
            resource_id=str(id),
            owners={"workflows": {"org_id": owner}},
        )
        result = await session.execute(select(Workflow).where(Workflow.id == id))
        wf = result.scalars().first()
        if not wf:
            raise HTTPException(status_code=404, detail="Workflow not found")
        rootResult = await resolve_root_id(wf.root_workflow_id, wf.id)
        versionResult = await next_version(wf.version)
        newWf = Workflow(org_id=wf.org_id, root_workflow_id=rootResult.root_workflow_id, title=wf.title, trigger_event=wf.trigger_event, version=versionResult.version)
        session.add(newWf)
        await session.flush()
        await session.execute(
            update(Action).values(new_workflow_id=newWf.id, source_workflow_id=wf.id)
        )
        return {
            "workflow": newWf,
        }


async def execute_with_report(session: AsyncSession, id: int, current_user: dict | None = None):
    async with session.begin():
        owner_row = await session.execute(
            select(Workflows.org_id).where(Workflows.id == id)
        )
        owner = owner_row.scalar_one_or_none()
        await authz_check(
            current_user,
            action="ExecuteWorkflow",
            resource="workflow",
            resource_id=str(id),
            owners={"workflows": {"org_id": owner}},
        )
        result = await session.execute(select(Workflow).where(Workflow.id == id))
        wf = result.scalars().first()
        if not wf:
            raise HTTPException(status_code=404, detail="Workflow not found")
        # @state workflows.ExecuteWorkflow — transition guard
        allowed_execute_workflow: dict[str, bool] = {
            "active": True,
        }
        if wf.status not in allowed_execute_workflow:
            raise HTTPException(status_code=409, detail="Workflow is not active")
        result = await session.execute(select(Organization).where(Organization.id == wf.id))
        org = result.scalars().first()
        if not org:
            raise HTTPException(status_code=404, detail="Organization not found")
        if is_zero_balance(org.credits_balance):
            raise HTTPException(status_code=402, detail="Insufficient credits")
        actionResult = await process_action(wf.trigger_event, params)
        await session.execute(
            update(Organization).where(Organization.id == org.id).values(amount=1)
        )
        reportResult = await generate_report(1, "completed", wf.title)
        log = ExecutionLog(credits_spent=1, org_id=wf.org_id, report_key=reportResult.report_key, status="completed", workflow_id=wf.workflow_id)
        session.add(log)
        await session.flush()
        return {
            "action_result": actionResult,
            "execution_log": log,
            "report": reportResult,
        }


async def execute_workflow(session: AsyncSession, id: int, current_user: dict | None = None, event_bus: EventBus | None = None):
    async with session.begin():
        owner_row = await session.execute(
            select(Workflows.org_id).where(Workflows.id == id)
        )
        owner = owner_row.scalar_one_or_none()
        await authz_check(
            current_user,
            action="ExecuteWorkflow",
            resource="workflow",
            resource_id=str(id),
            owners={"workflows": {"org_id": owner}},
        )
        result = await session.execute(select(Workflow).where(Workflow.id == id))
        wf = result.scalars().first()
        if not wf:
            raise HTTPException(status_code=404, detail="Workflow not found")
        # @state workflows.ExecuteWorkflow — transition guard
        allowed_execute_workflow: dict[str, bool] = {
            "active": True,
        }
        if wf.status not in allowed_execute_workflow:
            raise HTTPException(status_code=409, detail="Workflow is not active")
        result = await session.execute(select(Organization).where(Organization.id == wf.id))
        org = result.scalars().first()
        if not org:
            raise HTTPException(status_code=404, detail="Organization not found")
        if is_zero_balance(org.credits_balance):
            raise HTTPException(status_code=402, detail="Insufficient credits")
        actionResult = await process_action(wf.trigger_event, params)
        await session.execute(
            update(Organization).where(Organization.id == org.id).values(amount=1)
        )
        log = ExecutionLog(credits_spent=1, org_id=wf.org_id, status="completed", workflow_id=wf.workflow_id)
        session.add(log)
        await session.flush()
        await event_bus.publish("workflow.executed", {
            "OrgID": wf.org_id,
            "Status": "completed",
            "WorkflowID": wf.id,
        })
        return {
            "action_result": actionResult,
            "execution_log": log,
        }


async def get_workflow(session: AsyncSession, id: int, current_user: dict | None = None):
    owner_row = await session.execute(
        select(Workflows.org_id).where(Workflows.id == id)
    )
    owner = owner_row.scalar_one_or_none()
    await authz_check(
        current_user,
        action="GetWorkflow",
        resource="workflow",
        resource_id=str(id),
        owners={"workflows": {"org_id": owner}},
    )
    result = await session.execute(select(Workflow).where(Workflow.id == id))
    wf = result.scalars().first()
    if not wf:
        raise HTTPException(status_code=404, detail="Workflow not found")
    return {
        "workflow": wf,
    }


async def list_execution_logs(session: AsyncSession, id: int, current_user: dict | None = None):
    owner_row = await session.execute(
        select(Workflows.org_id).where(Workflows.id == id)
    )
    owner = owner_row.scalar_one_or_none()
    await authz_check(
        current_user,
        action="ListExecutionLogs",
        resource="workflow",
        resource_id=str(id),
        owners={"workflows": {"org_id": owner}},
    )
    result = await session.execute(select(Workflow).where(Workflow.id == id))
    wf = result.scalars().first()
    if not wf:
        raise HTTPException(status_code=404, detail="Workflow not found")
    result = await session.execute(select(ExecutionLog).where(ExecutionLog.workflow_id == wf.workflow_id))
    execution_logs = result.scalars().all()
    return {
        "execution_logs": execution_logs,
    }


async def list_workflow_versions(session: AsyncSession, id: int, current_user: dict | None = None):
    owner_row = await session.execute(
        select(Workflows.org_id).where(Workflows.id == id)
    )
    owner = owner_row.scalar_one_or_none()
    await authz_check(
        current_user,
        action="ListWorkflowVersions",
        resource="workflow",
        resource_id=str(id),
        owners={"workflows": {"org_id": owner}},
    )
    result = await session.execute(select(Workflow).where(Workflow.id == id))
    wf = result.scalars().first()
    if not wf:
        raise HTTPException(status_code=404, detail="Workflow not found")
    result = await session.execute(select(Workflow).where(Workflow.org_id == current_user["org_id"], Workflow.root_id == wf.id))
    workflows = result.scalars().all()
    return {
        "workflows": workflows,
    }


async def list_workflows(session: AsyncSession, current_user: dict | None = None):
    await authz_check(
        current_user,
        action="ListWorkflows",
        resource="workflow",
    )
    result = await session.execute(select(Workflow).where(Workflow.org_id == current_user["org_id"]))
    workflows = result.scalars().all()
    return {
        "workflows": workflows,
    }


async def pause_workflow(session: AsyncSession, id: int, current_user: dict | None = None):
    async with session.begin():
        owner_row = await session.execute(
            select(Workflows.org_id).where(Workflows.id == id)
        )
        owner = owner_row.scalar_one_or_none()
        await authz_check(
            current_user,
            action="PauseWorkflow",
            resource="workflow",
            resource_id=str(id),
            owners={"workflows": {"org_id": owner}},
        )
        result = await session.execute(select(Workflow).where(Workflow.id == id))
        wf = result.scalars().first()
        if not wf:
            raise HTTPException(status_code=404, detail="Workflow not found")
        # @state workflows.PauseWorkflow — transition guard
        allowed_pause_workflow: dict[str, bool] = {
            "active": True,
        }
        if wf.status not in allowed_pause_workflow:
            raise HTTPException(status_code=409, detail="Cannot pause workflow")
        await session.execute(
            update(Workflow).where(Workflow.id == wf.id).values(status="paused")
        )
        result = await session.execute(select(Workflow).where(Workflow.id == wf.id))
        updated = result.scalars().first()
        return {
            "workflow": updated,
        }


async def save_workflow_actions(session: AsyncSession, id: int, body: SaveWorkflowActionsRequest, current_user: dict | None = None):
    async with session.begin():
        owner_row = await session.execute(
            select(Workflows.org_id).where(Workflows.id == id)
        )
        owner = owner_row.scalar_one_or_none()
        await authz_check(
            current_user,
            action="SaveWorkflowActions",
            resource="workflow",
            resource_id=str(id),
            owners={"workflows": {"org_id": owner}},
        )
        result = await session.execute(select(Workflow).where(Workflow.id == id))
        wf = result.scalars().first()
        if not wf:
            raise HTTPException(status_code=404, detail="Workflow not found")
        await session.execute(delete(Action).where(Action.workflow_id == wf.workflow_id))
        await session.execute(
            update(Action).values(items=body.actions_json, workflow_id=wf.workflow_id)
        )
        return {
            "message": "ok",
        }


