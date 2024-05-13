# Asset Uploader Service

This project is a file uploader service using Go, Gin framework for the API, PostgreSQL for storing file metadata, and MinIO as the object storage solution. It provides endpoints for uploading files, listing all uploaded files, and downloading files.

## Run the project
- install and run docker
- git clone https://github.com/r2st/asset_service.git
- cd go_asset_service
- ./start.sh
- open in your browser: http://localhost:3000

## Prerequisites

- Docker and Docker Compose (for running PostgreSQL, MinIO, go_service and asser-uploader locally)
- Install docker from https://www.docker.com/products/docker-desktop/

## API Endpoints

- **POST /upload** (http://localhost:8080/upload): Upload a file. The file will be stored in MinIO, and its metadata will be saved in PostgreSQL.
- **GET /files** (http://localhost:8080/files): List all files along with their metadata.
- **GET /files** (http://localhost:8080/files?page=1&limit=100): List all files along with their metadata based on page and limit. "page" is value current page value and limit is the number of rows in each page.
- **GET /download/:id** (http://localhost:8080/download/db7fa1dc-3a29-46c3-840e-ac9db00aba75): Download a file by its ID.

## Featues (Completed)
- one HTTP endpoint that accepts, uploads, and stores assets in the cloud.
- one HTTP endpoint that allows retrieval of an uploaded asset.
- one HTTP endpoint that returns a list of all files in the system, their identifier, original filename, and the byte size of the file.
- A web page/app that provides a browser-based mechanism for using the upload, list and download endpoints.
- Automated the setup such that you could easily deploy a working copy of the app using a single command.

## Architecture Overview

- **API Server** : Built with the Gin framework, handles HTTP requests for uploading, listing, and downloading files.
- **Database** : PostgreSQL used for storing file metadata like original filename, storage filename, file extension, and file size.
- **Object Storage** : MinIO, used for storing the actual files uploaded through the API.
- **React Web App (Asset Uploader)** : A react JS based web app is developed to test the features and the end points.

## External Go Libraries

This project utilizes several external libraries to handle various functionalities like HTTP server management, database interactions, and object storage. Here's an overview of each library and its purpose:

### Gin-Gonic (Gin)

- **Imported As:** `github.com/gin-gonic/gin`
- **Purpose:** Gin is a web framework written in Go (Golang). It features a Martini-like API with much better performance, up to 40 times faster. It's used to create RESTful API endpoints and handle HTTP requests and responses more efficiently.
- **Usage:** In this project, Gin is used to route HTTP requests to the appropriate handlers, manage middleware, and process API responses.

### PostgreSQL Driver (pgx)

- **Imported As:** `github.com/jackc/pgx/v4/stdlib`
- **Purpose:** `pgx` is a PostgreSQL driver for Go. It provides support for advanced PostgreSQL features not available in the standard `database/sql` library. It's known for its high performance.
- **Usage:** Used in this project to connect to the PostgreSQL database and execute SQL queries efficiently with connection pooling capabilities.

### MinIO Go Client

- **Imported As:** `github.com/minio/minio-go/v7`
- **Purpose:** The MinIO Client SDK provides simple APIs to access any Amazon S3 compatible object storage server. 
- **Usage:** This client is used to interact with MinIO services, handling operations like uploading files, downloading files, and checking bucket existence.

### CORS Middleware

- **Imported As:** `github.com/gin-contrib/cors`
- **Purpose:** CORS middleware for Gin provides an easy-to-use API to manage Cross-Origin Resource Sharing (CORS), allowing or restricting resources based on the client's origin.
- **Usage:** It's used to configure CORS settings for the API, specifying which domains, methods, and headers are allowed.

### UUID

- **Imported As:** `github.com/google/uuid`
- **Purpose:** This package provides immutable UUIDs based on RFC 4122, which can be used to generate unique identifiers for objects without requiring a central authority.
- **Usage:** It is used in this project to generate unique identifiers for file storage, ensuring each file has a unique name.

### Standard Libraries

These are Go's core libraries used for basic functionality like HTTP server management, file path handling, and more:

#### `net/http`

- **Purpose:** Provides HTTP client and server implementations.
- **Usage:** Used to listen and serve HTTP requests in the application.

#### `database/sql`

- **Purpose:** Generic interface around SQL (or SQL-like) databases.
- **Usage:** Used to interact with the database using generic interface methods.

#### `os`

- **Purpose:** Provides a platform-independent interface to operating system functionality.
- **Usage:** Used to read environment variables and interact with the OS file system.

#### `log`

- **Purpose:** Provides a simple logging interface.
- **Usage:** Used for logging information about the application's operation, errors, and other important system messages.

#### `time`

- **Purpose:** Provides functionality for measuring and displaying time.
- **Usage:** Used to handle timeouts and timestamping events.

#### `context`

- **Purpose:** Provides mechanisms to carry deadlines, cancellation signals, and other request-scoped values across API boundaries and between processes.
- **Usage:** Used to manage request-scoped values, cancellation signals, and deadlines across API calls.

#### `encoding/json`

- **Purpose:** Implements encoding and decoding of JSON as defined in RFC 7159.
- **Usage:** Used to encode and decode data structures to/from JSON, especially used for Config.

#### `io`

- **Purpose:** Provides basic interfaces to I/O primitives.
- **Usage:** Used in file handling operations, especially to read and write data to HTTP responses and requests.

#### `path/filepath`

- **Purpose:** Implements utility routines for manipulating filename paths in a way compatible with the target operating system-defined file paths.
- **Usage:** Used to handle file paths during file operations, ensuring compatibility across different operating systems. Used here for extracting the file extension from file name.
