from pydantic import BaseModel
from typing import Optional


class AddActionRequest(BaseModel):
    action_type: str
    config: str
    sequence_order: int


class CreateWorkflowRequest(BaseModel):
    title: str
    trigger_event: str


class SaveWorkflowActionsRequest(BaseModel):
    actions_json: str


