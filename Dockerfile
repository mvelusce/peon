FROM golang:1.15 AS builder

WORKDIR /opt/peon/app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN cd cmd/peon && go build -o ../../bin/peon

FROM python:3.7
RUN pip install virtualenv

WORKDIR /opt/peon/app

COPY --from=builder /opt/peon/app/bin/peon /opt/peon/app/bin/

ENV PATH="/opt/peon/app/bin:${PATH}"

CMD ["./bin/peon"]
