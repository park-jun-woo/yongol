from fastapi import HTTPException
from sqlalchemy import select, update, delete
from sqlalchemy.ext.asyncio import AsyncSession

async def clone_template(session: AsyncSession, params: dict, body: dict, user: dict | None = None):
    async with session.begin():
        # @auth template.CloneTemplate
        # TODO: integrate OPA policy evaluation
        result = await session.execute(select(Template).where(Template.ID == request["id"]))
        tmpl = result.scalars().first()
        if not tmpl:
            raise HTTPException(status_code=404, detail="Template not found")
        result = await session.execute(select(Workflow).where(Workflow.ID == tmpl["SourceWorkflowID"]))
        sourceWf = result.scalars().first()
        if not sourceWf:
            raise HTTPException(status_code=404, detail="Source workflow not found")
        newWf = Workflow(OrgID=currentUser["OrgID"], Title=sourceWf["Title"], TriggerEvent=sourceWf["TriggerEvent"])
        session.add(newWf)
        await session.flush()
        await session.execute(
            update(Action).where(Action.id == params["id"]).values(**body)
        )
        await session.execute(
            update(Template).where(Template.id == params["id"]).values(**body)
        )
        return {
            "workflow": newWf,
        }


from fastapi import HTTPException
from sqlalchemy import select, update, delete
from sqlalchemy.ext.asyncio import AsyncSession

async def get_template(session: AsyncSession, params: dict, body: dict, user: dict | None = None):
    result = await session.execute(select(Template).where(Template.ID == request["id"]))
    template = result.scalars().first()
    if not template:
        raise HTTPException(status_code=404, detail="Template not found")
    return {
        "template": template,
    }


from fastapi import HTTPException
from sqlalchemy import select, update, delete
from sqlalchemy.ext.asyncio import AsyncSession

async def list_templates(session: AsyncSession, params: dict, body: dict, user: dict | None = None):
    result = await session.execute(select(Template).where(Template.Category == request["category"], Template.Cursor == request["cursor"], Template.PerPage == request["per_page"]))
    items = result.scalars().all()
    return {
        "items": items,
    }


from fastapi import HTTPException
from sqlalchemy import select, update, delete
from sqlalchemy.ext.asyncio import AsyncSession

async def publish_template(session: AsyncSession, params: dict, body: dict, user: dict | None = None):
    async with session.begin():
        # @auth template.PublishTemplate
        # TODO: integrate OPA policy evaluation
        result = await session.execute(select(Template).where(Template.SourceWorkflowID == request["source_workflow_id"]))
        existing = result.scalars().first()
        if existing:
            raise HTTPException(status_code=409, detail="Already published")
        template = Template(Category=request["category"], Description=request["description"], OrgID=currentUser["OrgID"], SourceWorkflowID=request["source_workflow_id"], Title=request["title"])
        session.add(template)
        await session.flush()
        return {
            "template": template,
        }


