import { Inject, Injectable, OnModuleInit } from '@nestjs/common';
import { ClientGrpc } from '@nestjs/microservices';
import { Observable } from 'rxjs';
import {
  GalleryClient,
  GALLERY_SERVICE_NAME,
  KONACHAN_V1_PACKAGE_NAME,
  PostsResponse,
} from '../generated/apis/konachan';

@Injectable()
export class KonachanService implements OnModuleInit {
  private galleryService: GalleryClient;

  constructor(@Inject(KONACHAN_V1_PACKAGE_NAME) private client: ClientGrpc) {}

  onModuleInit() {
    this.galleryService = this.client.getService<GalleryClient>(GALLERY_SERVICE_NAME);
  }

  getPosts(tags: string[]): Observable<PostsResponse> {
    return this.galleryService.posts({ tags });
  }
}
