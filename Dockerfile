FROM golang:1.22

RUN mkdir -p /opt/go
COPY . /opt/go
WORKDIR /opt/go
RUN go build -o bin/lls

ENTRYPOINT ["/opt/go/run.sh"]
