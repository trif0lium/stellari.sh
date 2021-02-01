package trif0lium.stellarish.nhentai;

import io.grpc.Server;
import io.grpc.ServerBuilder;
import java.io.IOException;

public class App 
{
    public static void main(String[] args) throws IOException, InterruptedException {
        Server server = ServerBuilder.forPort(50053).addService(new NHImpl()).build();
        server.start();
        server.awaitTermination();
    }
}
