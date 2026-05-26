//ff:func feature=gen-fastapi type=generator control=iteration dimension=1
//ff:what RenderMain — FastAPI main.py 부트스트랩 Python 소스 생성

package boot

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// RenderMain produces the FastAPI main.py bootstrap file content. It reads
// the BootPlan to determine which initialization blocks (CORS, auth,
// middleware, etc.) should be included.
func RenderMain(plan *ir.BootPlan, featureNames []string) (string, error) {
	if plan == nil {
		return "", fmt.Errorf("RenderMain: nil plan")
	}

	var b strings.Builder

	b.WriteString("from contextlib import asynccontextmanager\n")
	b.WriteString("from fastapi import FastAPI\n")

	if hasActiveBlock(plan, "cors") {
		b.WriteString("from fastapi.middleware.cors import CORSMiddleware\n")
	}

	b.WriteString("from app.database import engine, Base\n")

	// Import routers
	for _, fm := range featureNames {
		b.WriteString(fmt.Sprintf("from app.routers import %s as %s_router\n", fm, fm))
	}

	b.WriteString("\n\n")

	// Lifespan context manager
	b.WriteString("@asynccontextmanager\n")
	b.WriteString("async def lifespan(app: FastAPI):\n")
	b.WriteString("    async with engine.begin() as conn:\n")
	b.WriteString("        await conn.run_sync(Base.metadata.create_all)\n")
	b.WriteString("    yield\n")
	b.WriteString("    await engine.dispose()\n\n")

	// App creation
	b.WriteString(fmt.Sprintf("app = FastAPI(title=\"%s\", lifespan=lifespan)\n\n", plan.ProjectID))

	// CORS middleware
	if hasActiveBlock(plan, "cors") {
		b.WriteString("app.add_middleware(\n")
		b.WriteString("    CORSMiddleware,\n")
		b.WriteString("    allow_origins=[\"*\"],\n")
		b.WriteString("    allow_credentials=True,\n")
		b.WriteString("    allow_methods=[\"*\"],\n")
		b.WriteString("    allow_headers=[\"*\"],\n")
		b.WriteString(")\n\n")
	}

	// Include routers
	for _, fm := range featureNames {
		b.WriteString(fmt.Sprintf("app.include_router(%s_router.router)\n", fm))
	}
	b.WriteString("\n")

	// Health endpoint
	b.WriteString("\n@app.get(\"/health\")\n")
	b.WriteString("async def health():\n")
	b.WriteString("    return {\"status\": \"ok\"}\n")

	return b.String(), nil
}
