FROM golang:1.21

WORKDIR /node-discovery
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -o /usr/local/bin/node-discovery

CMD ["node-discovery"]