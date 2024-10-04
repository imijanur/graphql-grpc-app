# Define variables
BIN_DIR = ./bin
GQL_BINARY = $(BIN_DIR)/graphql_server
GRPC_BINARY = $(BIN_DIR)/grpc_server

# Directories
PROTO_DIR = .
PROTO_OUT_DIR = ./proto

# Protocol buffer compiler and flags
PROTOC = protoc
PROTOC_GEN_GO = protoc-gen-go
PROTOC_GEN_GRPC_GO = protoc-gen-go-grpc

# gRPC services
PROTO_FILES = $(wildcard *.proto)

# Go files to be generated
GO_OUT = --go_out=$(PROTO_OUT_DIR) --go_opt=paths=source_relative
GRPC_OUT = --go-grpc_out=$(PROTO_OUT_DIR) --go-grpc_opt=paths=source_relative

# Environment file for database config
ENV_FILE = .env

# Include environment variables
include $(ENV_FILE)
export $(shell sed 's/=.*//' $(ENV_FILE))

# Compile protobuf definitions
build-proto:
	sudo apt install protoc-gen-go
	sudo apt install protoc-gen-go-grpc
	$(PROTOC) -I=$(PROTO_DIR) $(GO_OUT) $(GRPC_OUT) $(PROTO_FILES)

# Build GraphQL and gRPC binaries
build-gql:
	GOOS=linux GOARCH=amd64 go build -o $(GQL_BINARY) ./cmd/graphql_server/*.go

build-grpc:
	GOOS=linux GOARCH=amd64 go build -o $(GRPC_BINARY) ./cmd/grpc_server/*.go

build-all: build-gql build-grpc

# Run services
run-gql: build-gql
	$(GQL_BINARY)

run-grpc: build-grpc
	$(GRPC_BINARY)

# Generate GraphQL Go code from schema
generate-gql:
	go get github.com/99designs/gqlgen
	go run github.com/99designs/gqlgen generate

# Generate SQLBoiler models
generate-sqlboiler:
	go get -u -t github.com/vattle/sqlboiler
	sqlboiler mysql

# Run GraphQL and gRPC code generation
generate-all: generate-gql generate-sqlboiler

generate-users:
	go run migrations/*.go

# Migrate, build & run all steps
run: run-migrations generate-sqlboiler build-grpc run-grpc

# Test
test:
	go test ./...

# Cleanup binaries and generated files
clean:
	rm -rf $(BIN_DIR) internal/gql/generated.go internal/gql/models_gen.go

