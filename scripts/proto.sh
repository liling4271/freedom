#!/bin/bash
# scripts/generate-proto.sh

set -e

echo "Cleaning generated files..."
rm -rf gen/proto
mkdir -p gen/proto

echo "Generating proto files..."

# 生成 common proto
protoc --proto_path=internal/shared/proto \
       --proto_path=internal/shared/third_party/googleapis \
       --go_out=gen/proto \
       --go_opt=paths=source_relative \
       internal/shared/proto/common/common.proto

# 生成 user proto
protoc --proto_path=internal/shared/proto \
       --proto_path=internal/shared/third_party/googleapis \
       --go_out=gen/proto \
       --go_opt=paths=source_relative \
       --go-grpc_out=gen/proto \
       --go-grpc_opt=paths=source_relative \
       --grpc-gateway_out=gen/proto \
       --grpc-gateway_opt=paths=source_relative \
       --grpc-gateway_opt=logtostderr=true \
       internal/shared/proto/user/message.proto \
       internal/shared/proto/user/user.proto \
       internal/shared/proto/user/enum.proto

echo "Proto files generated successfully!"
echo "Generated files are in: gen/proto"