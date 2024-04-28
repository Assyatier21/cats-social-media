FROM golang:1.19

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN apt-get update && apt-get install -y default-mysql-client

COPY tools/migration.sh /usr/local/bin/

RUN chmod 777 /usr/local/bin/migration.sh

RUN go build -o /app_bin ./cmd/main.go

EXPOSE 8800

ENTRYPOINT ["migration.sh"]

CMD ["/app_bin"]