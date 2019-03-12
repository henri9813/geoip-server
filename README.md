# geoip-server
 A simple web GeoIP server
 
# Usage

|     Database    |      Url      | Status |
|:---------------:|:-------------:|:------:|
|   GeoLite-City  |   /city/{IP}  |   ok   |
| GeoLite-Country | /country/{IP} |   ok   |
|   GeoLite-ASN   |   /asn/{IP}   |   ok   | 
       
```bash
user@machine $ curl http://geoipserver/asn/8.8.8.8
{"AutonomousSystemNumber":15169,"AutonomousSystemOrganization":"Google LLC"}
```

# Installation

## Download database

To automatically download MaxMind GeoIpLite you can use the makefile as follow:
```bash
make download
```

## You have golang >=1.12 installed

To run server, you just have to run the main.go
```bash
go run main.go
```

## You want to use docker

Please up the docker-compose.
```bash
docker-compose up -d
```
