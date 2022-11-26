FROM golang:1.19.2

WORKDIR /go/delivery

# We should probably only copy what's needed to make sure this doesn't exponentially grow
COPY . .

RUN go build -o recipe-tracker-ms

RUN groupadd --gid 1000 go && useradd --uid 1000 --gid go --shell /bin/bash --create-home go

USER go
EXPOSE 8080

CMD ./recipe-tracker-ms