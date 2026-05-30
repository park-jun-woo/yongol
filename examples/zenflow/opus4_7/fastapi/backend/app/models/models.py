from __future__ import annotations

import uuid
from datetime import date, datetime, timedelta
from decimal import Decimal
from typing import Any

from sqlalchemy import (
    Boolean, Date, DateTime, Float, Integer, Interval,
    LargeBinary, Numeric, String, Text, Uuid,
)
from sqlalchemy.dialects.postgresql import ARRAY, INET, JSONB
from sqlalchemy.orm import Mapped, mapped_column

from app.database import Base


class FullendQueue(Base):
    __tablename__ = "fullend_queue"

    id: Mapped[int] = mapped_column(Integer, primary_key=True)
    topic: Mapped[str] = mapped_column(Text)
    payload: Mapped[dict[str, Any]] = mapped_column(JSONB)
    priority: Mapped[str] = mapped_column(Text, default="normal")
    status: Mapped[str] = mapped_column(Text, default="pending")
    created_at: Mapped[datetime] = mapped_column(DateTime(timezone=True))
    deliver_at: Mapped[datetime] = mapped_column(DateTime(timezone=True))
    processed_at: Mapped[datetime | None] = mapped_column(DateTime(timezone=True), nullable=True)
    traceparent: Mapped[str] = mapped_column(Text)


class FullendSessions(Base):
    __tablename__ = "fullend_sessions"

    key: Mapped[str] = mapped_column(Text, primary_key=True)
    value: Mapped[bytes] = mapped_column(LargeBinary)
    expires_at: Mapped[datetime] = mapped_column(DateTime(timezone=True))


class Organizations(Base):
    __tablename__ = "organizations"

    id: Mapped[int] = mapped_column(Integer, primary_key=True)
    name: Mapped[str] = mapped_column(Text)
    plan_type: Mapped[str] = mapped_column(String, default="free")
    credits_balance: Mapped[int] = mapped_column(Integer)
    latitude: Mapped[str] = mapped_column(Text)
    longitude: Mapped[str] = mapped_column(Text)
    address_verified: Mapped[bool] = mapped_column(Boolean)
    created_at: Mapped[datetime] = mapped_column(DateTime(timezone=True))


class Templates(Base):
    __tablename__ = "templates"

    id: Mapped[int] = mapped_column(Integer, primary_key=True)
    source_workflow_id: Mapped[int] = mapped_column(Integer)
    org_id: Mapped[int] = mapped_column(Integer)
    title: Mapped[str] = mapped_column(Text)
    description: Mapped[str] = mapped_column(Text)
    category: Mapped[str] = mapped_column(String)
    clone_count: Mapped[int] = mapped_column(Integer)
    created_at: Mapped[datetime] = mapped_column(DateTime(timezone=True))


class Users(Base):
    __tablename__ = "users"

    id: Mapped[int] = mapped_column(Integer, primary_key=True)
    org_id: Mapped[int] = mapped_column(Integer)
    email: Mapped[str] = mapped_column(String)
    password_hash: Mapped[str] = mapped_column(String)
    role: Mapped[str] = mapped_column(String)
    claims: Mapped[dict[str, Any]] = mapped_column(JSONB, default="{}")
    created_at: Mapped[datetime] = mapped_column(DateTime(timezone=True))


class Actions(Base):
    __tablename__ = "actions"

    id: Mapped[int] = mapped_column(Integer, primary_key=True)
    workflow_id: Mapped[int] = mapped_column(Integer)
    action_type: Mapped[str] = mapped_column(Text)
    config: Mapped[str] = mapped_column(Text)
    sequence_order: Mapped[int] = mapped_column(Integer)


class RefreshTokens(Base):
    __tablename__ = "refresh_tokens"

    token_hash: Mapped[str] = mapped_column(Text, primary_key=True)
    claims: Mapped[dict[str, Any]] = mapped_column(JSONB)
    expires_at: Mapped[datetime] = mapped_column(DateTime(timezone=True))
    revoked_at: Mapped[datetime | None] = mapped_column(DateTime(timezone=True), nullable=True)
    created_at: Mapped[datetime] = mapped_column(DateTime(timezone=True))


class Webhooks(Base):
    __tablename__ = "webhooks"

    id: Mapped[int] = mapped_column(Integer, primary_key=True)
    org_id: Mapped[int] = mapped_column(Integer)
    url: Mapped[str] = mapped_column(Text)
    event_type: Mapped[str] = mapped_column(String)
    created_at: Mapped[datetime] = mapped_column(DateTime(timezone=True))


class Workflows(Base):
    __tablename__ = "workflows"

    id: Mapped[int] = mapped_column(Integer, primary_key=True)
    org_id: Mapped[int] = mapped_column(Integer)
    title: Mapped[str] = mapped_column(Text)
    trigger_event: Mapped[str] = mapped_column(Text)
    status: Mapped[str] = mapped_column(String, default="draft")
    version: Mapped[int] = mapped_column(Integer)
    root_workflow_id: Mapped[int] = mapped_column(Integer)
    assigned_to: Mapped[int] = mapped_column(Integer)
    assignment_confidence: Mapped[str] = mapped_column(String, default="none")
    created_at: Mapped[datetime] = mapped_column(DateTime(timezone=True))


class AuditLogs(Base):
    __tablename__ = "audit_logs"

    id: Mapped[int] = mapped_column(Integer, primary_key=True)
    org_id: Mapped[int] = mapped_column(Integer)
    actor_id: Mapped[int] = mapped_column(Integer)
    action: Mapped[str] = mapped_column(String)
    resource_type: Mapped[str] = mapped_column(String)
    resource_id: Mapped[str] = mapped_column(Text)
    detail: Mapped[str] = mapped_column(Text)
    created_at: Mapped[datetime] = mapped_column(DateTime(timezone=True))


class ExecutionLogs(Base):
    __tablename__ = "execution_logs"

    id: Mapped[int] = mapped_column(Integer, primary_key=True)
    workflow_id: Mapped[int] = mapped_column(Integer)
    org_id: Mapped[int] = mapped_column(Integer)
    status: Mapped[str] = mapped_column(String)
    credits_spent: Mapped[int] = mapped_column(Integer)
    report_key: Mapped[str] = mapped_column(String)
    executed_at: Mapped[datetime] = mapped_column(DateTime(timezone=True))


class FullendCache(Base):
    __tablename__ = "fullend_cache"

    key: Mapped[str] = mapped_column(Text, primary_key=True)
    value: Mapped[bytes] = mapped_column(LargeBinary)
    expires_at: Mapped[datetime] = mapped_column(DateTime(timezone=True))

