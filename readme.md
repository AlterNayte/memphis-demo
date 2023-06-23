export PATH="$PATH:$(go env GOPATH)/bin"

protoc --go_out=. ./test.proto 