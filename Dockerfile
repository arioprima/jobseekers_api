FROM golang:1.22.6-alpine

RUN mkdir /build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN mkdir -p /build

RUN  go build -o jobseeker-api cmd/main.go

COPY jobseeker-api /build/

COPY config.yaml /build/

COPY docs /build/

WORKDIR /build

CMD ["./jobseeker-api"]
