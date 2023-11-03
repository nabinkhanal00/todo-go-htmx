FROM golang:1.19

WORKDIR /output

COPY . /output

RUN apt-get update

RUN apt-get install sqlite3

RUN go mod tidy

RUN go build -o main cmd/app/main.go

ENV PORT 8080

EXPOSE 8080

CMD [ /output/main ]
