# Simple REST-like communication protocol

## Messages types overview
There are two types of messages:
- Requests
- Responses

Requests are sent by client and can be:
- Authentication request
- Status request
- Start scanning request
- Stop scanning request
- Send scanned data from time to time (default today)

Responses are sent by server and can be:
- Authentication ok
- Authentication not ok
- Server status:
 - Free RAM
 - Free disk
 - Database connection status
 - Scanning status
- Scanning started
- Scanning stopped
- Scanned data

## Implementation details

Packed into JSON. JSON representation TBD.
