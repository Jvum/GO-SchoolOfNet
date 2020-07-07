FROM golang:1.8-alpine

WORKDIR /go/src/gotesting
ADD ./src/gotesting /go/src/gotesting

RUN go install .

CMD ["gotesting"]