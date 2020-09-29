FROM golang:1.15

WORKDIR /opt/peon/app

COPY . .

RUN go get -d ./...
RUN cd cmd/peon && go build -o ../../bin/peon

RUN rm -rf cmd/ internal/ test/ tools/

ENV PATH="/opt/peon/app/bin:${PATH}"

CMD ["./bin/peon"]
