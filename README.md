## TAXI

This project contains simple http server, that stores 50 random orders (two random a-z symbols)

At the start of the application, 50 random orders are being generated.
Every 200ms, a random order is canceled and new one appears.

Default port is `8080`

The REST API consists of two links:

1. `/request` by which an imaginary taxi driver receives a random order
from 50 existing in the system.
2. `/admin/requests` by which the administrator receives information about all 
created (and canceled) requests in the system and how many times each of the requests 
have been shown to taxi drivers (zero is omitted).

### Dependencies

This project uses go modules for vendoring. Run `make vendor` to download dependencies.


### Run

To start the application user `go build taxi && ./taxi` 
