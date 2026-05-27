"""Authorization stub — integrate OPA/Rego policy evaluation."""
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
