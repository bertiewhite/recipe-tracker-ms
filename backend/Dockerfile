FROM golang:1.19-alpine AS build

WORKDIR /app

# We should probably only copy what's needed to make sure this doesn't exponentially grow
COPY . .

RUN go build -o /app/service

FROM alpine:3.17

COPY --from=build /app /app

EXPOSE 8080

ENTRYPOINT ["/app/service"]