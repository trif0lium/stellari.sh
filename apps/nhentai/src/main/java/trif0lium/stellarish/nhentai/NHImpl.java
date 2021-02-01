package trif0lium.stellarish.nhentai;

import io.grpc.stub.StreamObserver;

public class NHImpl extends NHGrpc.NHImplBase {
    @Override
    public void gallery(GalleryRequest request, StreamObserver<GalleryResponse> responseObserver) {
        responseObserver.onNext(
                GalleryResponse
                        .newBuilder()
                        .setId(-1)
                        .setMediaId(-1)
                        .setTitle(
                                GalleryResponse.Title
                                        .newBuilder()
                                        .setEnglish("")
                                        .setJapanese("")
                                        .setPretty("")
                                        .build()
                        )
                        .build()
        );
        responseObserver.onCompleted();
    }
}
