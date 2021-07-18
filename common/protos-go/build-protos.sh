#!/bin/bash

BUILDDIR=./
PROTOSDIR=../protos
FABRIC_PROTOSDIR=../fabric-protos

protoc --proto_path=$PROTOSDIR --proto_path=$FABRIC_PROTOSDIR --go_out=plugins=grpc:$BUILDDIR --go_opt=paths=source_relative $PROTOSDIR/common/query.proto $PROTOSDIR/common/ack.proto $PROTOSDIR/common/proofs.proto $PROTOSDIR/common/state.proto $PROTOSDIR/common/access_control.proto $PROTOSDIR/common/membership.proto $PROTOSDIR/common/verification_policy.proto $PROTOSDIR/common/interop_payload.proto $PROTOSDIR/common/asset_locks.proto
protoc --proto_path=$PROTOSDIR --proto_path=$FABRIC_PROTOSDIR --go_out=plugins=grpc:$BUILDDIR --go_opt=paths=source_relative $PROTOSDIR/fabric/view_data.proto
protoc --proto_path=$PROTOSDIR --proto_path=$FABRIC_PROTOSDIR --go_out=plugins=grpc:$BUILDDIR --go_opt=paths=source_relative $PROTOSDIR/corda/view_data.proto
protoc --proto_path=$PROTOSDIR --proto_path=$FABRIC_PROTOSDIR --go_out=plugins=grpc:$BUILDDIR --go_opt=paths=source_relative $PROTOSDIR/networks/networks.proto
protoc --proto_path=$PROTOSDIR --proto_path=$FABRIC_PROTOSDIR --go_out=plugins=grpc:$BUILDDIR --go_opt=paths=source_relative $PROTOSDIR/relay/datatransfer.proto
protoc --proto_path=$PROTOSDIR --proto_path=$FABRIC_PROTOSDIR --go_out=plugins=grpc:$BUILDDIR --go_opt=paths=source_relative $PROTOSDIR/driver/driver.proto
#protoc --proto_path=$PROTOSDIR --proto_path=$FABRIC_PROTOSDIR --go_out=plugins=grpc --go_out=$BUILDDIR --go_opt=paths=source_relative $PROTOSDIR/common/query.proto $PROTOSDIR/common/ack.proto $PROTOSDIR/common/proofs.proto $PROTOSDIR/common/state.proto $PROTOSDIR/corda/view_data.proto $PROTOSDIR/fabric/view_data.proto $PROTOSDIR/common/access_control.proto $PROTOSDIR/common/membership.proto $PROTOSDIR/common/verification_policy.proto $PROTOSDIR/common/interop_payload.proto $PROTOSDIR/common/asset_locks.proto
