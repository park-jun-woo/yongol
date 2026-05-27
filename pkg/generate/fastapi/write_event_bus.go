//ff:func feature=gen-fastapi type=util control=sequence
//ff:what writeEventBus — FastAPI EventBus stub 파일 기록

package fastapi

import (
	"os"
	"path/filepath"
)

// writeEventBus writes the event bus dependency stub module.
func writeEventBus(appDir string) error {
	depsDir := filepath.Join(appDir, "dependencies")
	if err := os.MkdirAll(depsDir, 0o755); err != nil {
		return err
	}

	content := `"""Event bus stub — replace with your preferred message broker."""
import logging

logger = logging.getLogger(__name__)


class EventBus:
    """In-memory event bus stub. Replace with Redis Pub/Sub, RabbitMQ, etc."""

    async def publish(self, topic: str, payload: dict) -> None:
        logger.info("publish %s: %s", topic, payload)
        # TODO: integrate with your message broker


event_bus = EventBus()
`
	return os.WriteFile(filepath.Join(depsDir, "event_bus.py"), []byte(content), 0o644)
}
