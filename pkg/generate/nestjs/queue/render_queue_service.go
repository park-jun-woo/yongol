//ff:func feature=gen-nestjs type=generator control=sequence
//ff:what RenderQueueService — QueueService TypeScript 소스 생성

package queue

// RenderQueueService produces the NestJS queue service file content.
// The stub provides a publish method that downstream users implement
// with their preferred message broker (BullMQ, RabbitMQ, etc.).
func RenderQueueService() string {
	return `import { Injectable, Logger } from '@nestjs/common';

@Injectable()
export class QueueService {
  private readonly logger = new Logger(QueueService.name);

  async publish(topic: string, payload: Record<string, any>): Promise<void> {
    this.logger.log(` + "`" + `publish ${topic}: ${JSON.stringify(payload)}` + "`" + `);
    // TODO: integrate with your message broker (BullMQ, RabbitMQ, etc.)
  }
}
`
}
