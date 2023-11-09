Certainly! Here's the provided markdown converted into a proper README.md:

```markdown
# Lightning-Fast URL Shortener with Golang Fiber and Redis

A high-performance URL shortener built with Golang using the Fiber web framework and Redis as the database.

## Features

- Lightning-fast URL shortening.
- Utilizes the Golang Fiber framework for efficient handling.
- Uses Redis for quick and reliable storage.

## Getting Started

### Prerequisites

- Golang: [Install Golang](https://golang.org/doc/install)
- Redis: [Install Redis](https://redis.io/download)

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/lightning-fast-url-shortener.git
   ```

2. Change into the project directory:

   ```bash
   cd lightning-fast-url-shortener
   ```

3. Build and run the project:

   ```bash
   go build
   ./url-shortener
   ```
   
## API Endpoint

### Shorten URL

- **URL:** `POST /`
- **Request:**
  - **Method:** `POST`
  - **Endpoint:** `/`
  - **Request Body:**
    ```json
    {
      "url": "https://www.yourdomain.com"
    }
    ```
- **Response (Success):**
  - **Status Code:** `200 OK`
  - **Response Body:**
    ```json
    {
      "shorted_url": "http://localhost:3000/s/abc123",
      "error": ""
    }
    ```
- **Response (Error):**
  - **Status Code:** `400 Bad Request`
  - **Response Body:**
    ```json
    {
      "shorted_url": "",
      "error": "Invalid URL"
    }
    ```
```
