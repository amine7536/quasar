FROM golang:latest

COPY gobgp.toml /app/.
RUN go get github.com/osrg/gobgp/...
