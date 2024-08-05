# Simple Go HTTP Server

This project implements a basic HTTP server in Go using only the "net" package. It's designed as a learning tool to understand the fundamentals of HTTP servers without relying on higher-level libraries.

## Features

- Handles GET request
- Simple routing mechanism for "/" and "/user" paths
- 404 Not Found response for undefined routes
- Processes connections sequentially (non-concurrent)

## Limitations

- Limited to handling one connection at a time
- Basic HTTP parsing without full protocol compliance
- Fixed buffer size (1 kb) may not handle large requests efficiently
- Lacks support for other HTTP methods (POST, PUT, DELETE, etc.)
- No support for HTTPS

## Further Improvements

1. Implement concurrent connection handling using goroutines
2. Add support for more HTTP methods
3. Improve request parsing to handle headers and larger bodies more efficiently
4. Implement a more robust routing system
5. Add HTTPS support
6. Enhance error handling and add comprehensive logging
7. Implement middleware support for common functionalities (e.g., authentication)
8. Add configuration options (e.g., port number, timeout settings)

## Testing

To test the server:

1. Run the server:
  ```
  go run main.go
  ```
2. Use curl or a web browser to test GET requests:
  ```
  curl http://localhost:8080/
  curl http://localhost:8080/user
  ```
3. Test undefined routes for 404 responses:
  ```
  curl http//localhost:8080/asd
  ``` 
