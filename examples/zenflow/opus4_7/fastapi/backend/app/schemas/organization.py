from pydantic import BaseModel
from typing import Optional


class VerifyOrgAddressRequest(BaseModel):
    address: str


