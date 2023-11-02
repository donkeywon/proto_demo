protoc -I=. --go-grpc_out=pb --go_out=pb --validate_out="lang=go:pb" proto/*.proto
