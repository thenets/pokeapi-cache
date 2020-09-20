# Build
FROM golang:1.15-alpine
WORKDIR $GOPATH/src/github.com/thenets/pokeapi-cache 
RUN apk add git
COPY . .
RUN go get -d -v ./...
RUN go build -o /tmp/pokeapi-cache
RUN chmod +x /tmp/pokeapi-cache

# Server
FROM alpine
ENV PORT=8080
RUN adduser -S -D -H -h /app/ pikachu
USER pikachu
WORKDIR /app/
COPY --from=0 /tmp/pokeapi-cache  ./
CMD ["./pokeapi-cache"]
