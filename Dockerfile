FROM golang:1.19 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./
COPY validator ./validator

RUN CGO_ENABLED=0 GOOS=linux go build -o /creditcard-validator


FROM golang:1.19

ENV GIN_MODE=release
ENV PORT=3005

COPY --from=builder /creditcard-validator /creditcard-validator

EXPOSE $PORT

CMD ["/creditcard-validator"]
