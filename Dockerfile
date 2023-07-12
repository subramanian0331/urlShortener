FROM golang:1.20
LABEL authors="subbu"

WORKDIR /app
COPY . /app

RUN go build -o /go-url-shortener


EXPOSE 8000

#ENTRYPOINT ["top", "-b"]
CMD ["/go-url-shortener"]