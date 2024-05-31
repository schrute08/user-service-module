# User Service Module

This repository contains a Golang gRPC service for managing user details with search functionality. The service maintains a list of user details and provides endpoints to fetch user details by ID, list user details by a list of IDs, and search user details based on criteria like city, phone number, and marital status.

## Project Structure
```bash
    .
    ├── cmd
    │ ├── client
    │ │ └── main.go
    │ └── server
    │ └── main.go
    ├── internal
    │ ├── errors
    │ │ └── errors.go
    │ ├── server
    │ │ ├── user_server.go
    │ │ └── user_server_test.go
    │ └── utils
    │ ├── validations.go
    │ └── validations_test.go
    └── proto
    └── user
    ├── user.proto
    └── userpb
    ├── user.pb.go
    └── user_grpc.pb.go
```
- **cmd**: Contains client and server applications entry points.
- **internal**: Holds internal package code.
  - **errors**: Defines custom error types.
  - **server**: Implements gRPC server and its tests.
  - **utils**: Provides utility functions for validation and testing.
- **proto**: Contains protocol buffer definitions.
  - **user**: Protobuf definition files for user service.
    - **user.proto**: Protobuf file defining user service API.
    - **userpb**: Generated Go code for user service.

## Features

- Mocked database by storing a list of User structs.
- Fetch user details by user ID.
- Fetch user details list by a list of user IDs.
- Search user details based on city, phone number, and marital status.

## Prerequisites

- Go (1.16 or later)
- Protocol Buffers compiler (`protoc`)
- Docker (for containerization)
- Basic understanding of gRPC and Protocol Buffers, you can refer to my blog [here](https://medium.com/@schrute08/mastering-grpc-building-a-go-based-microservice-architecture-cbbba70e52f5).

## Setup and Installation

1. **Clone the repository:**

```bash
git clone https://github.com/yourusername/grpc-user-service.git
cd grpc-user-service
```

2. **Install Go dependencies:**
    
```bash
go mod download
```

## Build and Run the Application
### Running Locally
Run the application, you will need to run the server and client separately in two different terminals:

```
go run cmd/server/main.go
go run cmd/client/main.go
```
### Running in Docker
To build Docker images for the client and server, use the provided Dockerfiles in their respective directories.

```
docker build -t user-service-client ./cmd/client
docker build -t user-service-server ./cmd/server
```

Then, you can run the containers:

```
docker run -d --name=user-service-server user-service-server
docker run -d --name=user-service-client user-service-client
```

## Testing
To run the tests, and see the coverage, execute the following command:

```bash
go test ./... -cover
```