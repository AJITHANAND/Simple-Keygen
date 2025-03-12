# Go Random Key API

This project is a simple REST API server written in Go that generates a random key every day. Users can retrieve the current key without any authentication.

## Project Structure

```
go-rest-api
├── src
│   ├── main.go          # Entry point of the application
│   ├── handlers
│   │   └── keyHandler.go # Handles GET requests for the current key
│   └── utils
│       └── keyGenerator.go # Generates and rotates the random key
├── Dockerfile           # Dockerfile to build the application image
├── go.mod               # Module definition and dependencies
└── README.md            # Project documentation
```

## Features

- Generates a random key every 24 hours.
- Provides an endpoint to retrieve the current key.
- No authentication required for accessing the key.

## Getting Started

### Prerequisites

- Go (1.16 or later)
- Docker (optional, for containerization)

### Running the Application

1. Clone the repository:

   ```
   git clone https://github.com/AJITHANAND/Simple-Keygen.git simple-keygen
   cd simple-keygen
   ```

2. Install dependencies:

   ```
   go mod tidy
   ```

3. Run the application:

   ```
   go run src/main.go
   ```

   The server will start on `localhost:8080`.

### Accessing the Key

You can retrieve the current random key by sending a GET request to:

```
http://localhost:8080/key
```

### Running with Docker

To build and run the application in a Docker container:

1. Build the Docker image:

   ```
   docker build -t go-random-key-api .
   ```

2. Run the Docker container:

   ```
   docker run -p 8080:8080 go-random-key-api
   ```

The server will be accessible at `http://localhost:8080/key`.

## License

This project is licensed under the MIT License. See the LICENSE file for details.
