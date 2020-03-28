FROM golang:1.13.8
RUN mkdir /api_service
ADD . /api_service

WORKDIR /api_service
RUN go build -o main .
CMD ["/api_service/main"]

EXPOSE 8080