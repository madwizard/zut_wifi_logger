# gowifiscanner
WiFiScanner in go fo learning purposes

## Arch

* Main program is a standalone server
* Server runs as a daemon and collects data
* Server makes data available over socket file and HTTPS
* Server transmits data using JSON format
* Most basic client is command line utility communicating over socket file
* Most basic client doesn't use any graphical library

## Database

* Data will be kept either in json files or sqlite
* DB structure is described in file DB.md


