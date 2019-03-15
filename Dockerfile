FROM golang AS build-env
ADD . /src

RUN cd /src && make download && CGO_ENABLED=0 GOOS=linux go build -o geoip-server

FROM alpine
WORKDIR /app
COPY --from=build-env /src/geoip-server /app/
COPY --from=build-env /src/data /app/data

ENTRYPOINT ./geoip-server
