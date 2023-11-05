FROM golang:1.21-alpine AS go
WORKDIR /app
COPY go.mod go.sum main.go ./
RUN go mod download \
  && go build -o main /app/main.go

FROM alpine
WORKDIR /app
COPY --from=go /app/main .
CMD [ "/app/main" ]