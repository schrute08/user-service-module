# gRPC User Service Module

This repository contains a Golang gRPC service for managing user details with search functionality. The service maintains a list of user details and provides endpoints to fetch user details by ID, list user details by a list of IDs, and search user details based on criteria like city, phone number, and marital status.

## Features

- Fetch user details by user ID.
- Mocked database by storing a list of User structs.
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
1. Build the application:

    ```bash
        go build -o main .
    ```
2. Run the application:

    ```bash
        ./main
    ```

