cd $(dirname $0)
mkdir -p generated
protoc \
  --go_out=./generated \
  --go-grpc_out=./generated \
  ./apis/danbooru.proto \
  --go_opt=paths=source_relative \
  --go-grpc_opt=paths=source_relative
