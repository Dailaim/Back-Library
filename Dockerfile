FROM docker.io/golang:1.20

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .
RUN go build -v -o .
CMD ["./back-library"]