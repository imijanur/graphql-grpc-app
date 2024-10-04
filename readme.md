# GraphQL and gRPC Server

This project is a GraphQL and gRPC server built with Go. It uses SQLBoiler for ORM and gqlgen for GraphQL schema generation.

## Prerequisites

Before you begin, ensure you have met the following requirements:

- **Go**: You need to have Go installed. You can download it from [here](https://golang.org/dl/).
- **Make**: Ensure you have `make` installed. This is typically available on Unix-based systems. For Windows, you can use [GnuWin](http://gnuwin32.sourceforge.net/packages/make.htm).
- **MySQL**: You need to have MySQL installed and running. You can download it from [here](https://dev.mysql.com/downloads/mysql/).
- **SQLBoiler**: SQLBoiler is used for ORM. Install it using the following command:
  ```sh
  go install github.com/volatiletech/sqlboiler/v4@latest
  go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@latest
  ```

## Installing
1. Clone the repository:
   ```sh
   git clone https://github.com/imijanur/graphql-grpc-server.git
   cd graphql-grpc-server
   ```

2. Create a `.env` file in the root directory and add the following environment variables:
   ```sh
   DB_USER=root
   DB_PASSWORD=password
   DB_NAME=graphql_grpc_server
   DB_HOST=localhost
   DB_PORT=33061
   ```
   Replace the values with your MySQL credentials.

3. compile proto files:
   ```sh
   make build-proto
   ```
   This will generate go codes related to grpc

4. Build Graphql schemas and server:
   ```sh
   make build-gql
   ```
   This generate graphql server binary

5. build Grpc server:
   ```sh
   make build-grpc
   ```
   This will generate the grpc server binary.

6. Build both server :
   ```sh
   make build-all
   ```
   This will generate the GraphQL and gRPC servers.

7. Run Graphql server :
   ```sh
   make run-gql
   ```
   This will start the GraphQL server.

8. Run GRPC server :
   ```sh
   make run-grpc
   ```
   This will start the GRPC server.

9. Generate graphql :
   ```sh
   make generate-gql
   ```
   This will generate graphql related files based on the schema.

9. Generate SQLBOILER:
   ```sh
   make generate-sqlboiler
   ```
   This will generate sqlboiler models

10. Generate Users:
   ```sh
   make generate-user
   ```
   This will generate 20000 dummy users and related data

11. test:
   ```sh
   make test
   ```
   this will test all available test files



## Usage

1. GraphQL Playground: Open the GraphQL Playground in your browser by visiting `http://localhost:8080/`.

2. gRPC Server: The gRPC server is running on port `50051`.

