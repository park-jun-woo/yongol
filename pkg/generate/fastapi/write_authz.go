//ff:func feature=gen-fastapi type=util control=sequence
//ff:what writeAuthz — FastAPI authz stub 파일 기록

package fastapi

import (
	"os"
	"path/filepath"
)

// writeAuthz writes the authorization dependency stub module.
func writeAuthz(appDir string) error {
	depsDir := filepath.Join(appDir, "dependencies")
	if err := os.MkdirAll(depsDir, 0o755); err != nil {
		return err
	}

	content := `"""Authorization stub — integrate OPA/Rego policy evaluation."""
import logging

from fastapi import HTTPException

logger = logging.getLogger(__name__)


async def authz_check(
    user: dict | None,
    action: str,
    resource: str,
    **context,
) -> None:
    """Evaluate an authorization policy.

    Replace with actual OPA/Rego evaluation. Raises HTTPException(403) when
    the policy rejects the request.
    """
    logger.debug("authz check: %s on %s (user=%s)", action, resource, user)
    # TODO: integrate OPA/Rego policy evaluation
    # raise HTTPException(status_code=403, detail="access denied")
`
	return os.WriteFile(filepath.Join(depsDir, "authz.py"), []byte(content), 0o644)
}
