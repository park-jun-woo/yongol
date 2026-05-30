import {
  Controller,
  Get,
  Param,
  Body,
  Req,
} from '@nestjs/common';
import { GetDashboardService } from './getDashboard.service';

@Controller('dashboard')
export class GetDashboardController {
  constructor(private readonly service: GetDashboardService) {}

  @Get('')
  async getDashboard(
    @Req() req: any,
    @Param() params: any,
    @Body() body: any,
  ) {
    return this.service.getDashboard(params, body, req.user);
  }
}
