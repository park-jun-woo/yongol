import { Injectable, Logger } from '@nestjs/common';

@Injectable()
export class QueueService {
  private readonly logger = new Logger(QueueService.name);

  async publish(topic: string, payload: Record<string, any>): Promise<void> {
    this.logger.log(`publish ${topic}: ${JSON.stringify(payload)}`);
    // TODO: integrate with your message broker (BullMQ, RabbitMQ, etc.)
  }
}
