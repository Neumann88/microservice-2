#!/usr/bin/env bash

DIR="./api/protobuf"
GEN="./internal/gen/proto"

rm -rf ${GEN}
mkdir ${GEN}

for dir in $(find ${DIR} -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq); do
    files=$(find "${dir}" -name '*.proto')

    protoc -I ${DIR} \
        --go_out ${GEN} --go_opt paths=source_relative \
        --go-grpc_out ${GEN} --go-grpc_opt paths=source_relative \
        "${files}"
done
