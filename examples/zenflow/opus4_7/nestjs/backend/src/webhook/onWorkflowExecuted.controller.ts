import {
  Controller,
} from '@nestjs/common';
import { OnWorkflowExecutedService } from './onWorkflowExecuted.service';

@Controller('webhook')
export class OnWorkflowExecutedController {
  constructor(private readonly service: OnWorkflowExecutedService) {}

  // Subscribe handler for topic: workflow.executed
  async handleOnWorkflowExecuted(payload: any) {
    return this.service.onWorkflowExecuted(payload);
  }
}
