from pydantic import BaseModel
from typing import Optional


class CreateWebhookRequest(BaseModel):
    event_type: str
    url: str


