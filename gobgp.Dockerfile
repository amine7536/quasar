FROM golang:latest

COPY gobgp.toml /gobgp.toml
RUN go get github.com/osrg/gobgp/...
