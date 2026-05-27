import {
  Controller,
  Get,
  Req,
} from '@nestjs/common';
import { GetDashboardService } from './getDashboard.service';

@Controller('dashboard')
export class GetDashboardController {
  constructor(private readonly service: GetDashboardService) {}

  @Get('')
  async getDashboard(
    @Req() req: any,
  ) {
    return this.service.getDashboard(req.user);
  }
}
