# DomainWatch API

DomainWatch is an API developed in Go (version 1.24) aimed at providing various domain-related functionalities. This project is under development, and many features are yet to be implemented.

## Current Features

- **WHOIS Lookup**: Retrieve detailed information about a domain's registration.
- **Domain Validation**: Check if a domain is valid.
- **Active TLS Versions**: Identify which TLS versions are active on a domain.

## Technologies Used

- **Language**: Go 1.24
- **Framework**: Gin Gonic (for route management and HTTP server)
- **Tools**: `openssl` and `curl` for TLS checks.

## How to Run

1. Ensure you have Go 1.24 installed on your machine.
2. Clone this repository:
   ```bash
   git clone https://github.com/Mamonha/domainwatch.git
   ```
3. Navigate to the project directory:
   ```bash
   cd domainwatch
   ```
4. Run the project:
   ```bash
   go run main.go
   ```
5. Access the API at `http://localhost:8080`.

## Usage Examples

### WHOIS Lookup

**Endpoint**: `GET /api/domain/whois/:domain`

**Request Example**:
```bash
curl http://localhost:8080/api/domain/whois/example.com
```

### Domain Validation

**Endpoint**: `GET /api/domain/validate/:domain`

**Request Example**:
```bash
curl http://localhost:8080/api/domain/validate/example.com
```

### TLS Check

**Endpoint**: `POST /api/domain/tls`

**Request Example**:
```bash
curl -X POST -H "Content-Type: application/json" -d '{"domain": "example.com"}' http://localhost:8080/api/domain/tls
```

## Future Features

- DNS verification.
- Get Certificates(b64, pem);
- Support for more security protocols.
- Integration with third-party services for domain analysis.
- Performance and scalability improvements.

## Contribution

Contributions are welcome! Feel free to open issues and submit pull requests.

## License

This project is licensed under the MIT license.