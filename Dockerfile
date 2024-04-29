# syntax=docker/dockerfile:1

FROM golang:1.22-alpine3.19 AS build

WORKDIR /app

COPY . .

RUN adduser --disabled-password \
  --gecos "" \
  --home "/home/nohome" \
  --shell "/sbin/nologin" \
  --no-create-home \
  --uid 1001 nologinuser

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /gobin ./src

FROM scratch

WORKDIR /

COPY --from=build /gobin /gobin

COPY --from=build /etc/passwd /etc/passwd

USER nologinuser

EXPOSE 8080

ENTRYPOINT ["/gobin"]