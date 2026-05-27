from pydantic import BaseModel
from typing import Optional


class PublishTemplateRequest(BaseModel):
    category: str
    description: str
    source_workflow_id: int
    title: str


