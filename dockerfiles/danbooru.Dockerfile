FROM golang:latest as build-env
WORKDIR /app/danbooru
ADD ./apps/danbooru /app/danbooru
ADD ./apis /apis
RUN go build main.go

FROM gcr.io/distroless/base
COPY --from=build-env /app/danbooru/main /
CMD ["./main"]
