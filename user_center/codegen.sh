proto=$1
protoc --go_out=plugins=grpc:. $proto