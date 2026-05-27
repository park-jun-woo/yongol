from pydantic import BaseModel
from typing import Optional


class SetScheduleRequest(BaseModel):
    cron: str


