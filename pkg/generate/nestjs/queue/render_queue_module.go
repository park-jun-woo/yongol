//ff:func feature=gen-nestjs type=generator control=sequence
//ff:what RenderQueueModule — QueueModule TypeScript 소스 생성

package queue

// RenderQueueModule produces the NestJS queue module file content.
func RenderQueueModule() string {
	return `import { Global, Module } from '@nestjs/common';
import { QueueService } from './queue.service';

@Global()
@Module({
  providers: [QueueService],
  exports: [QueueService],
})
export class QueueModule {}
`
}
