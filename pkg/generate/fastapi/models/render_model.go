//ff:func feature=gen-fastapi type=generator control=iteration dimension=1
//ff:what RenderModels — DDL Table 배열 → SQLAlchemy 2.0 Mapped model Python 소스 생성

package models

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// RenderModels produces a SQLAlchemy 2.0 models file from DDL tables.
// Each table maps to a class inheriting from Base with Mapped columns.
func RenderModels(tables []ddl.Table) (string, error) {
	var b strings.Builder

	b.WriteString("from __future__ import annotations\n\n")
	b.WriteString("import uuid\n")
	b.WriteString("from datetime import date, datetime, timedelta\n")
	b.WriteString("from decimal import Decimal\n")
	b.WriteString("from typing import Any\n\n")
	b.WriteString("from sqlalchemy import (\n")
	b.WriteString("    Boolean, Date, DateTime, Float, ForeignKey, Index, Integer, Interval,\n")
	b.WriteString("    LargeBinary, Numeric, String, Text, UniqueConstraint, Uuid,\n")
	b.WriteString(")\n")
	b.WriteString("from sqlalchemy.dialects.postgresql import ARRAY, INET, JSONB\n")
	b.WriteString("from sqlalchemy.orm import Mapped, mapped_column\n\n")
	b.WriteString("from app.database import Base\n")

	for _, table := range tables {
		b.WriteString("\n\n")
		if err := renderOneModel(&b, table); err != nil {
			return "", fmt.Errorf("RenderModels(%s): %w", table.Name, err)
		}
	}

	b.WriteString("\n")
	return b.String(), nil
}
