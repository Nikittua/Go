FROM golang:1.22.5-alpine3.19


WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go build -o /Go-Hello

EXPOSE 8080

CMD [ "/Go-Hello" ]
