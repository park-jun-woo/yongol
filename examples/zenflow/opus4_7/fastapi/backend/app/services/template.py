from fastapi import HTTPException
from sqlalchemy import select, update
from sqlalchemy.ext.asyncio import AsyncSession
from app.models.models import Action, Template, Workflow
from app.dependencies.authz import authz_check

async def clone_template(session: AsyncSession, id: int, current_user: dict | None = None):
    async with session.begin():
        owner_row = await session.execute(
            select(Templates.org_id).where(Templates.id == id)
        )
        owner = owner_row.scalar_one_or_none()
        await authz_check(
            current_user,
            action="CloneTemplate",
            resource="template",
            resource_id=id,
            resource_id=str(id),
            owners={"templates": {"org_id": owner}},
        )
        result = await session.execute(select(Template).where(Template.id == id))
        tmpl = result.scalars().first()
        if not tmpl:
            raise HTTPException(status_code=404, detail="Template not found")
        result = await session.execute(select(Workflow).where(Workflow.id == tmpl.id))
        sourceWf = result.scalars().first()
        if not sourceWf:
            raise HTTPException(status_code=404, detail="Source workflow not found")
        newWf = Workflow(org_id=current_user["org_id"], title=sourceWf.title, trigger_event=sourceWf.trigger_event)
        session.add(newWf)
        await session.flush()
        await session.execute(
            update(Action).values(new_workflow_id=newWf.id, source_workflow_id=sourceWf.id)
        )
        await session.execute(
            update(Template).where(Template.id == tmpl.id).values(**body)
        )
        return {
            "workflow": newWf,
        }


async def get_template(session: AsyncSession, id: int, current_user: dict | None = None):
    result = await session.execute(select(Template).where(Template.id == id))
    template = result.scalars().first()
    if not template:
        raise HTTPException(status_code=404, detail="Template not found")
    return {
        "template": template,
    }


async def list_templates(session: AsyncSession, category: str | None = None, cursor: str | None = None, per_page: int | None = None, current_user: dict | None = None):
    result = await session.execute(select(Template).where(Template.category == category).limit(per_page))
    items = result.scalars().all()
    return {
        "items": items,
    }


async def publish_template(session: AsyncSession, body: PublishTemplateRequest, current_user: dict | None = None):
    async with session.begin():
        await authz_check(
            current_user,
            action="PublishTemplate",
            resource="template",
        )
        result = await session.execute(select(Template).where(Template.source_workflow_id == body.source_workflow_id))
        existing = result.scalars().first()
        if existing:
            raise HTTPException(status_code=409, detail="Already published")
        template = Template(category=body.category, description=body.description, org_id=current_user["org_id"], source_workflow_id=body.source_workflow_id, title=body.title)
        session.add(template)
        await session.flush()
        return {
            "template": template,
        }


