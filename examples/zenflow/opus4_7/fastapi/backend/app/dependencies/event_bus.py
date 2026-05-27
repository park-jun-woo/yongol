"""Event bus stub — replace with your preferred message broker."""
import logging

logger = logging.getLogger(__name__)


class EventBus:
    """In-memory event bus stub. Replace with Redis Pub/Sub, RabbitMQ, etc."""

    async def publish(self, topic: str, payload: dict) -> None:
        logger.info("publish %s: %s", topic, payload)
        # TODO: integrate with your message broker


event_bus = EventBus()
