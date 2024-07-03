# URL Shortener Service in Go

## Project Overview

This project implements a simple URL shortener service using Go. It allows users to shorten long URLs and redirect to the actual URLs based on short codes.

struct Url
ID(unique)
Actual URL-
Short URL-actual url :short gareko creation date
CreationDate
for this project i am storing in the run time memory 


### Functionality

1. **Endpoints**:

   - `POST /shorturl`: Accepts a JSON payload with a long URL and returns a short URL.
   - `GET /redirect/{shortCode}`: Redirects users to the original long URL based on the short code provided.

2. **Data Storage**:
   - Uses an in-memory `map` to store mappings between short URLs and long URLs.

## Setup and Installation

### Prerequisites

- Go (version 1.13 or higher recommended)

### Dependencies

- `fmt`: Used for formatted I/O operations.
- `crypto/md5`: Used for MD5 hashing.
- `encoding/hex`: Used for hexadecimal encoding of hash values.
- `encoding/json`: Used for encoding and decoding JSON data.
- `errors`: Used for error handling.
- `net/http`: Used for creating HTTP servers and handling HTTP requests.
