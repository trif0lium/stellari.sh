import { Inject, Injectable, OnModuleInit } from '@nestjs/common';
import { ClientGrpc } from '@nestjs/microservices';
import { Observable } from 'rxjs';
import { GalleryResponse, NHClient, N_H_SERVICE_NAME } from 'src/generated/apis/nhentai';

@Injectable()
export class NhentaiService implements OnModuleInit {
  private nh: NHClient;

  constructor(@Inject(N_H_SERVICE_NAME) private client: ClientGrpc) {}

  onModuleInit() {
    this.nh = this.client.getService<NHClient>(N_H_SERVICE_NAME);
  }

  getGallery(id: number): Observable<GalleryResponse> {
    return this.nh.gallery({ id });
  }
}
