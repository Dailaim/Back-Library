FROM registry.fedoraproject.org/fedora-minimal:34

RUN microdnf install -y go git golang-x-net-devel

RUN microdnf update golang

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod tidy

RUN go mod download golang.org/x/net@v0.8.0

RUN go mod download


COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]