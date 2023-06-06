Install Protoc

`go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`
`go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`

Compile Protobuf

`protoc --proto_path=./protos ./protos/\*.proto --go_out=\$GOPATH/src --go-grpc_out=$GOPATH/src`

Execute

`go run main.go`
