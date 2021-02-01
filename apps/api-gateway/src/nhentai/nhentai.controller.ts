import { Controller, Get, Param, ParseIntPipe } from '@nestjs/common';
import { NhentaiService } from './nhentai.service';

@Controller('nhentai')
export class NhentaiController {
  constructor(private readonly nhentaiService: NhentaiService) {}

  @Get('/:id')
  getGallery(@Param('id', new ParseIntPipe()) id: number) {
    return this.nhentaiService.getGallery(id);
  }
}
