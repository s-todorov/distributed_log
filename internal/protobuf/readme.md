go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go: downloading google.golang.org/protobuf v1.27.1

go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go: downloading google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.2.0

go: downloading google.golang.org/grpc v1.44.0

export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin

protoc --go-grpc_out=. *.proto

protoc --go_out=. --go_opt=paths=source_relative     --go-grpc_out=. --go-grpc_opt=paths=source_relative     index.proto