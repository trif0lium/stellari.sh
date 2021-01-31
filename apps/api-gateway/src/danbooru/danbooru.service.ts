import { Inject, Injectable, OnModuleInit } from '@nestjs/common';
import { ClientGrpc } from '@nestjs/microservices';
import { Observable } from 'rxjs';
import {
  DANBOORU_V1_PACKAGE_NAME,
  GalleryClient,
  GALLERY_SERVICE_NAME,
  PostsResponse,
} from 'src/generated/apis/danbooru';

@Injectable()
export class DanbooruService implements OnModuleInit {
  private galleryService: GalleryClient;

  constructor(@Inject(DANBOORU_V1_PACKAGE_NAME) private client: ClientGrpc) {}

  onModuleInit() {
    this.galleryService = this.client.getService<GalleryClient>(GALLERY_SERVICE_NAME);
  }

  getPosts(tags: string[]): Observable<PostsResponse> {
    return this.galleryService.posts({ tags });
  }
}
