# ------------------- BUILD STAGE ------------------- #

FROM golang:1.19-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o job

# ------------------- DEV ------------------- #

FROM builder AS dev

WORKDIR /app

RUN apk add --no-cache make

RUN go install github.com/cespare/reflex@latest

CMD ["make watch"]

# ------------------- PROD ------------------- #

FROM alpine:latest AS prod

WORKDIR /

COPY --from=builder /app/job   /

CMD [ "./job" ]