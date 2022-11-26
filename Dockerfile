FROM golang:1.19.3

# We should probably only copy what's needed to make sure this doesn't exponentially grow
COPY . .

RUN go build -o recipe-tracker-ms

USER go
EXPOSE 8080

CMD ./recipe-tracker-ms