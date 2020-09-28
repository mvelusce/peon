FROM golang:1.14

WORKDIR /opt/peon/app

COPY . .

RUN go get -d ./...
RUN cd cmd/peon && go build -o ../../bin/peon

ENV PATH="/opt/peon/app/bin:${PATH}"

CMD ["./bin/peon"]
