FROM golang:1.18.1-alpine


WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go build -o /Go-Hello

EXPOSE 8080

CMD [ "/Go-Hello" ]
