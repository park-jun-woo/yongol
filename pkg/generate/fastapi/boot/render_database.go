//ff:func feature=gen-fastapi type=generator control=sequence
//ff:what RenderDatabase — FastAPI database.py (async SQLAlchemy engine + session) Python 소스 생성

package boot

import (
	"strings"
)

// RenderDatabase produces the database.py content with async SQLAlchemy
// engine, session maker, and Base declarative class.
func RenderDatabase() (string, error) {
	var b strings.Builder

	b.WriteString("from sqlalchemy.ext.asyncio import AsyncSession, create_async_engine\n")
	b.WriteString("from sqlalchemy.orm import DeclarativeBase, sessionmaker\n")
	b.WriteString("from app.config import settings\n\n")
	b.WriteString("engine = create_async_engine(settings.database_url, echo=settings.debug)\n\n")
	b.WriteString("async_session = sessionmaker(\n")
	b.WriteString("    engine, class_=AsyncSession, expire_on_commit=False\n")
	b.WriteString(")\n\n\n")
	b.WriteString("class Base(DeclarativeBase):\n")
	b.WriteString("    pass\n")

	return b.String(), nil
}
