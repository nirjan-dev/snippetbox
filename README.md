# Snippetbox

Snippetbox is a web application for pasting and sharing text snippets, built using Go. It serves as a pastebin-like tool where users can save and view code snippets or notes.

## About

This project was created while reading and following along with the **[Let's Go](https://lets-go.software/)** book by Alex Edwards. It implements the concepts and best practices taught in the book for building secure and scalable web applications with Go.

## Features

- **Secure Web Server**: Runs over HTTPS using TLS.
- **MySQL Database**: Uses MySQL for persistent storage of snippets and user data.
- **Session Management**: Secure, encrypted session handling.
- **User Authentication**: Secure signup, login, and logout functionality using bcrypt for password hashing.
- **Middleware Chains**: Implements standard middleware for security headers, logging, and panic recovery.
- **CSRF Protection**: Protects forms against Cross-Site Request Forgery attacks.
- **HTML Templating**: Server-side rendering with Go's `html/template` package.

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) (version 1.24+)
- [Docker Compose](https://docs.docker.com/compose/install/) (for running the database)

### Installation

1.  **Clone the repository** (if you haven't already).

2.  **Start the Database**:
    Use Docker Compose to start the MySQL database.
    ```bash
    docker-compose up -d
    ```
    This will start a MySQL instance on port `3306` with the credentials defined in `docker-compose.yml`.

3.  **Generate TLS Certificates**:
    The application runs on HTTPS and requires a TLS certificate and private key.
    ```bash
    mkdir -p tls
    cd tls
    go run $(go env GOROOT)/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost
    cd ..
    ```
    This generates `cert.pem` and `key.pem` in the `tls` directory.

4.  **Run the Application**:
    ```bash
    go run ./cmd/web
    ```

    You can also specify configuration via flags:
    ```bash
    go run ./cmd/web -addr=":4000" -dsn="web:password@/snippetbox?parseTime=true"
    ```

5.  **Access the App**:
    Open your browser and navigate to `https://localhost:4000`.
    
    *Note: Since this uses a self-signed certificate, your browser will display a security warning. You can safely bypass this for local development.*

## Project Structure

- `cmd/web`: Contains the application entry point (`main.go`) and web-specific code (handlers, middleware, routes).
- `pkg/models`: Contains the data models and database logic.
- `ui`: Contains the HTML templates and static files.
- `tls`: Stores the TLS certificates (ignored by git).