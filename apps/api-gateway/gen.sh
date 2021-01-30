cd $(dirname $0)
mkdir -p src/generated
protoc --plugin=./node_modules/.bin/protoc-gen-ts_proto \
  --ts_proto_out=./src/generated \
  ./apis/konachan.proto \
  --ts_proto_opt=nestJs=true
