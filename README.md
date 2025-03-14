# Go PlayGround

This is a simple Go application that starts an HTTP server and responds with a greeting message.

## Getting Started

### Prerequisites

- Go 1.16 or later

### Running the Application

To run the application, use the following command:

```sh
go run main.go
```

The server will start on port 8080.

Usage
You can access the greeting endpoint by navigating to: http://localhost:8080/hello?name=YourName

Replace YourName with your actual name to receive a personalized greeting.

Example
http://localhost:8080/hello?name=John

Response:
```bash
Hello John
```

### Accessing Swagger Documentation

To access the Swagger documentation for the API, navigate to the following URL in your web browser:

http://localhost:8080/swagger/index.html

This will display the Swagger UI, allowing you to interact with and test the API endpoints.

### Setting Up Swagger

To set up Swagger for your project, follow these steps:

1. **Install the `swag` CLI tool**:
   ```sh
   go install github.com/swaggo/swag/cmd/swag@latest
   export PATH=$PATH:$(go env GOPATH)/bin
   source ~/.zshrc
   ```

2. **Install the necessary packages**:
   ```sh
   go get github.com/swaggo/http-swagger
   ```

3. **Generate the Swagger documentation**:
   ```sh
   swag init
   ```
