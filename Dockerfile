FROM golang:1.18 as builder

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o server cmd/*.go

FROM golang:1.18.2-alpine 
WORKDIR /app

COPY --from=builder /app/server ./

EXPOSE 8080

