proto=$1
prefix=$2
protoc -I $prefix --go_out=plugins=grpc:./protos/gencode $proto