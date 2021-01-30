FROM rust:latest as build-env
WORKDIR /app/konachan
ADD ./apps/konachan /app/konachan
ADD ./apis /apis
RUN rustup component add rustfmt
RUN cargo build --release

FROM gcr.io/distroless/cc
COPY --from=build-env /app/konachan/target/release/konachan /
CMD ["./konachan"]
