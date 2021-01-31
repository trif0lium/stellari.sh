import { Controller, Get, ParseArrayPipe, Query } from '@nestjs/common';
import { Observable } from 'rxjs';
import { PostsResponse } from 'src/generated/apis/konachan';
import { KonachanService } from './konachan.service';

@Controller('konachan')
export class KonachanController {
  constructor(private readonly konachanService: KonachanService) {}

  @Get('/')
  getPosts(
    @Query('tags', new ParseArrayPipe({ items: String, separator: ',' })) tags: string[],
  ): Observable<PostsResponse> {
    return this.konachanService.getPosts(tags);
  }
}
