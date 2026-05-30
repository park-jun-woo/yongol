from fastapi import HTTPException
from sqlalchemy import select, update, delete
from sqlalchemy.ext.asyncio import AsyncSession

async def activate_workflow(session: AsyncSession, params: dict, body: dict, user: dict | None = None):
    async with session.begin():
        # @auth workflow.ActivateWorkflow
        # TODO: integrate OPA policy evaluation
        result = await session.execute(select(Workflow).where(Workflow.ID == request["id"]))
        wf = result.scalars().first()
        if not wf:
            raise HTTPException(status_code=404, detail="Workflow not found")
        result = await session.execute(select(Organization).where(Organization.ID == wf["OrgID"]))
        org = result.scalars().first()
        if not org:
            raise HTTPException(status_code=404, detail="Organization not found")
        if billing.is_zero_balance(org["CreditsBalance"]):
            raise HTTPException(status_code=402, detail="Insufficient credits")
        # @state workflows.ActivateWorkflow
        # TODO: integrate state machine validation
        await session.execute(
            update(Workflow).where(Workflow.id == params["id"]).values(**body)
        )
        result = await session.execute(select(Workflow).where(Workflow.ID == wf["ID"]))
        updated = result.scalars().first()
        return {
            "workflow": updated,
        }


from fastapi import HTTPException
from sqlalchemy import select, update, delete
from sqlalchemy.ext.asyncio import AsyncSession

async def add_action(session: AsyncSession, params: dict, body: dict, user: dict | None = None):
    async with session.begin():
        # @auth workflow.AddAction
        # TODO: integrate OPA policy evaluation
        result = await session.execute(select(Workflow).where(Workflow.ID == request["id"]))
        wf = result.scalars().first()
        if not wf:
            raise HTTPException(status_code=404, detail="Workflow not found")
        action = Action(ActionType=request["action_type"], Config=request["config"], SequenceOrder=request["sequence_order"], WorkflowID=wf["ID"])
        session.add(action)
        await session.flush()
        return {
            "action": action,
        }


from fastapi import HTTPException
from sqlalchemy import select, update, delete
from sqlalchemy.ext.asyncio import AsyncSession

async def archive_workflow(session: AsyncSession, params: dict, body: dict, user: dict | None = None):
    async with session.begin():
        # @auth workflow.ArchiveWorkflow
        # TODO: integrate OPA policy evaluation
        result = await session.execute(select(Workflow).where(Workflow.ID == request["id"]))
        wf = result.scalars().first()
        if not wf:
            raise HTTPException(status_code=404, detail="Workflow not found")
        # @state workflows.ArchiveWorkflow
        # TODO: integrate state machine validation
        await session.execute(
            update(Workflow).where(Workflow.id == params["id"]).values(**body)
        )
        result = await session.execute(select(Workflow).where(Workflow.ID == wf["ID"]))
        updated = result.scalars().first()
        return {
            "workflow": updated,
        }


from fastapi import HTTPException
from sqlalchemy import select, update, delete
from sqlalchemy.ext.asyncio import AsyncSession

async def auto_assign_workflow(session: AsyncSession, params: dict, body: dict, user: dict | None = None):
    async with session.begin():
        # @auth workflow.AutoAssignWorkflow
        # TODO: integrate OPA policy evaluation
        result = await session.execute(select(Workflow).where(Workflow.ID == request["id"]))
        wf = result.scalars().first()
        if not wf:
            raise HTTPException(status_code=404, detail="Workflow not found")
        result = await session.execute(select(User).where(User.OrgID == wf["OrgID"]))
        memberCount = result.scalars().first()
        match = await workflow.match_member(memberCount, wf["TriggerEvent"])
        await session.execute(
            update(Workflow).where(Workflow.id == params["id"]).values(**body)
        )
        result = await session.execute(select(Workflow).where(Workflow.ID == wf["ID"]))
        updated = result.scalars().first()
        if not updated:
            raise HTTPException(status_code=404, detail="Workflow not found")
        return {
            "workflow": updated,
        }


from fastapi import HTTPException
from sqlalchemy import select, update, delete
from sqlalchemy.ext.asyncio import AsyncSession

async def create_workflow(session: AsyncSession, params: dict, body: dict, user: dict | None = None):
    async with session.begin():
        # @auth workflow.CreateWorkflow
        # TODO: integrate OPA policy evaluation
        wf = Workflow(OrgID=currentUser["OrgID"], Title=request["title"], TriggerEvent=request["trigger_event"])
        session.add(wf)
        await session.flush()
        return {
            "workflow": wf,
        }


from fastapi import HTTPException
from sqlalchemy import select, update, delete
from sqlalchemy.ext.asyncio import AsyncSession

async def create_workflow_version(session: AsyncSession, params: dict, body: dict, user: dict | None = None):
    async with session.begin():
        # @auth workflow.CreateWorkflowVersion
        # TODO: integrate OPA policy evaluation
        result = await session.execute(select(Workflow).where(Workflow.ID == request["id"]))
        wf = result.scalars().first()
        if not wf:
            raise HTTPException(status_code=404, detail="Workflow not found")
        rootResult = await version.resolve_root_i_d(wf["RootWorkflowID"], wf["ID"])
        versionResult = await version.next_version(wf["Version"])
        newWf = Workflow(OrgID=wf["OrgID"], RootWorkflowID=rootResult["RootID"], Title=wf["Title"], TriggerEvent=wf["TriggerEvent"], Version=versionResult["Version"])
        session.add(newWf)
        await session.flush()
        await session.execute(
            update(Action).where(Action.id == params["id"]).values(**body)
        )
        return {
            "workflow": newWf,
        }


from fastapi import HTTPException
from sqlalchemy import select, update, delete
from sqlalchemy.ext.asyncio import AsyncSession

async def execute_with_report(session: AsyncSession, params: dict, body: dict, user: dict | None = None):
    async with session.begin():
        # @auth workflow.ExecuteWorkflow
        # TODO: integrate OPA policy evaluation
        result = await session.execute(select(Workflow).where(Workflow.ID == request["id"]))
        wf = result.scalars().first()
        if not wf:
            raise HTTPException(status_code=404, detail="Workflow not found")
        # @state workflows.ExecuteWorkflow
        # TODO: integrate state machine validation
        result = await session.execute(select(Organization).where(Organization.ID == wf["OrgID"]))
        org = result.scalars().first()
        if not org:
            raise HTTPException(status_code=404, detail="Organization not found")
        if billing.is_zero_balance(org["CreditsBalance"]):
            raise HTTPException(status_code=402, detail="Insufficient credits")
        actionResult = await worker.process_action(wf["TriggerEvent"], params)
        await session.execute(
            update(Organization).where(Organization.id == params["id"]).values(**body)
        )
        reportResult = await report.generate_report(1, "completed", wf["Title"])
        log = ExecutionLog(CreditsSpent=1, OrgID=wf["OrgID"], ReportKey=reportResult["ReportKey"], Status="completed", WorkflowID=wf["ID"])
        session.add(log)
        await session.flush()
        return {
            "action_result": actionResult,
            "execution_log": log,
            "report": reportResult,
        }


from fastapi import HTTPException
from sqlalchemy import select, update, delete
from sqlalchemy.ext.asyncio import AsyncSession
from app.dependencies.event_bus import EventBus

async def execute_workflow(session: AsyncSession, params: dict, body: dict, user: dict | None = None):
    async with session.begin():
        # @auth workflow.ExecuteWorkflow
        # TODO: integrate OPA policy evaluation
        result = await session.execute(select(Workflow).where(Workflow.ID == request["id"]))
        wf = result.scalars().first()
        if not wf:
            raise HTTPException(status_code=404, detail="Workflow not found")
        # @state workflows.ExecuteWorkflow
        # TODO: integrate state machine validation
        result = await session.execute(select(Organization).where(Organization.ID == wf["OrgID"]))
        org = result.scalars().first()
        if not org:
            raise HTTPException(status_code=404, detail="Organization not found")
        if billing.is_zero_balance(org["CreditsBalance"]):
            raise HTTPException(status_code=402, detail="Insufficient credits")
        actionResult = await worker.process_action(wf["TriggerEvent"], params)
        await session.execute(
            update(Organization).where(Organization.id == params["id"]).values(**body)
        )
        log = ExecutionLog(CreditsSpent=1, OrgID=wf["OrgID"], Status="completed", WorkflowID=wf["ID"])
        session.add(log)
        await session.flush()
        await event_bus.publish("workflow.executed", {
            "OrgID": wf["OrgID"],
            "Status": "completed",
            "WorkflowID": wf["ID"],
        })
        return {
            "action_result": actionResult,
            "execution_log": log,
        }


from fastapi import HTTPException
from sqlalchemy import select, update, delete
from sqlalchemy.ext.asyncio import AsyncSession

async def get_workflow(session: AsyncSession, params: dict, body: dict, user: dict | None = None):
    # @auth workflow.GetWorkflow
    # TODO: integrate OPA policy evaluation
    result = await session.execute(select(Workflow).where(Workflow.ID == request["id"]))
    wf = result.scalars().first()
    if not wf:
        raise HTTPException(status_code=404, detail="Workflow not found")
    return {
        "workflow": wf,
    }


from fastapi import HTTPException
from sqlalchemy import select, update, delete
from sqlalchemy.ext.asyncio import AsyncSession

async def list_execution_logs(session: AsyncSession, params: dict, body: dict, user: dict | None = None):
    # @auth workflow.ListExecutionLogs
    # TODO: integrate OPA policy evaluation
    result = await session.execute(select(Workflow).where(Workflow.ID == request["id"]))
    wf = result.scalars().first()
    if not wf:
        raise HTTPException(status_code=404, detail="Workflow not found")
    result = await session.execute(select(ExecutionLog).where(ExecutionLog.WorkflowID == wf["ID"]))
    execution_logs = result.scalars().all()
    return {
        "execution_logs": execution_logs,
    }


from fastapi import HTTPException
from sqlalchemy import select, update, delete
from sqlalchemy.ext.asyncio import AsyncSession

async def list_workflow_versions(session: AsyncSession, params: dict, body: dict, user: dict | None = None):
    # @auth workflow.ListWorkflowVersions
    # TODO: integrate OPA policy evaluation
    result = await session.execute(select(Workflow).where(Workflow.ID == request["id"]))
    wf = result.scalars().first()
    if not wf:
        raise HTTPException(status_code=404, detail="Workflow not found")
    result = await session.execute(select(Workflow).where(Workflow.OrgID == currentUser["OrgID"], Workflow.RootID == wf["ID"]))
    workflows = result.scalars().all()
    return {
        "workflows": workflows,
    }


from fastapi import HTTPException
from sqlalchemy import select, update, delete
from sqlalchemy.ext.asyncio import AsyncSession

async def list_workflows(session: AsyncSession, params: dict, body: dict, user: dict | None = None):
    # @auth workflow.ListWorkflows
    # TODO: integrate OPA policy evaluation
    result = await session.execute(select(Workflow).where(Workflow.OrgID == currentUser["OrgID"]))
    workflows = result.scalars().all()
    return {
        "workflows": workflows,
    }


from fastapi import HTTPException
from sqlalchemy import select, update, delete
from sqlalchemy.ext.asyncio import AsyncSession

async def pause_workflow(session: AsyncSession, params: dict, body: dict, user: dict | None = None):
    async with session.begin():
        # @auth workflow.PauseWorkflow
        # TODO: integrate OPA policy evaluation
        result = await session.execute(select(Workflow).where(Workflow.ID == request["id"]))
        wf = result.scalars().first()
        if not wf:
            raise HTTPException(status_code=404, detail="Workflow not found")
        # @state workflows.PauseWorkflow
        # TODO: integrate state machine validation
        await session.execute(
            update(Workflow).where(Workflow.id == params["id"]).values(**body)
        )
        result = await session.execute(select(Workflow).where(Workflow.ID == wf["ID"]))
        updated = result.scalars().first()
        return {
            "workflow": updated,
        }


from fastapi import HTTPException
from sqlalchemy import select, update, delete
from sqlalchemy.ext.asyncio import AsyncSession

async def save_workflow_actions(session: AsyncSession, params: dict, body: dict, user: dict | None = None):
    async with session.begin():
        # @auth workflow.SaveWorkflowActions
        # TODO: integrate OPA policy evaluation
        result = await session.execute(select(Workflow).where(Workflow.ID == request["id"]))
        wf = result.scalars().first()
        if not wf:
            raise HTTPException(status_code=404, detail="Workflow not found")
        await session.execute(delete(Action).where(Action.WorkflowID == wf["ID"]))
        await session.execute(
            update(Action).where(Action.id == params["id"]).values(**body)
        )
        return {
            "message": "ok",
        }


