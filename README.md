[![GitHub release](https://img.shields.io/github/release/igolaizola/h2-stream.svg)](https://github.com/igolaizola/h2-stream/releases)
[![Build Status](https://travis-ci.com/igolaizola/h2-stream.svg?branch=master)](https://travis-ci.com/igolaizola/h2-stream)
[![Go Report Card](https://goreportcard.com/badge/igolaizola/h2-stream)](http://goreportcard.com/report/igolaizola/h2-stream)
[![license](https://img.shields.io/github/license/igolaizola/h2-stream.svg)](https://github.com/igolaizola/h2-stream/blob/master/LICENSE.md)

# h2-stream

HTTP2 client and server implementation in GO that holds a persistent data stream
- Client takes data from standard input and forwards it to the server
- Client forwards server responses to standard output
- Server responds echoing the received data

TLS and non TLS options are both available

## Usage

### without TLS

launch server:
```
go run cmd/h2-stream/main.go serve --addr=localhost:8080 --tls=false
```

launch client:
```
go run cmd/h2-stream/main.go cli --addr=http://localhost:8080 --method=POST --data="BODY DATA"
```

### with TLS
launch server:
```
go run cmd/h2-stream/main.go serve --addr=localhost:8080 --tls=true --cert=certs/cert.pem --key=certs/key.pem
```

launch client:
```
go run cmd/h2-stream/main.go cli --addr=https://localhost:8080 --method=POST --data="BODY DATA" --insecure
```