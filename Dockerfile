FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go build -o /toronto-wheels ./cmd/toronto-wheels

EXPOSE 80

CMD [ "/toronto-wheels" ]
