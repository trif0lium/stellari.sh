FROM maven:latest as build-env
WORKDIR /app/nhentai
ADD ./apps/nhentai /app/nhentai
ADD ./apis /apis
RUN mvn clean compile package

FROM gcr.io/distroless/java
COPY --from=build-env /app/nhentai/target/nhentai-1.0-SNAPSHOT-jar-with-dependencies.jar /
CMD ["./nhentai-1.0-SNAPSHOT-jar-with-dependencies.jar"]
