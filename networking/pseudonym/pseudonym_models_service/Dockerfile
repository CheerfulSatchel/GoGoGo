FROM golang:1.13.8
RUN mkdir /models_service
ADD . /models_service

WORKDIR /models_service
RUN go build -o main .
CMD ["/models_service/main"]

EXPOSE 8081