FROM golang:1.22.6-alpine

RUN mkdir /build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN mkdir -p /build

RUN  go build -o jobseeker-api cmd/main.go

RUN cp jobseeker-api /build/

RUN cp config.yaml /build/

RUN cp -r docs /build/

WORKDIR /build

EXPOSE 8081

CMD ["./jobseeker-api"]
