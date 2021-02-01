import { Inject, Injectable, OnModuleInit } from '@nestjs/common';
import { ClientGrpc } from '@nestjs/microservices';
import { Observable } from 'rxjs';
import { GalleryResponse, NHClient, NHENTAI_V1_PACKAGE_NAME, N_H_SERVICE_NAME } from 'src/generated/apis/nhentai';

@Injectable()
export class NhentaiService implements OnModuleInit {
  private nh: NHClient;

  constructor(@Inject(NHENTAI_V1_PACKAGE_NAME) private client: ClientGrpc) {}

  onModuleInit() {
    this.nh = this.client.getService<NHClient>(N_H_SERVICE_NAME);
  }

  getGallery(id: number): Observable<GalleryResponse> {
    return this.nh.gallery({ id });
  }
}
