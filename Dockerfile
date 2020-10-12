FROM golang:1.14

WORKDIR /dnsbl-exporter
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
EXPOSE 8881

ENTRYPOINT ["dnsbl-exporter"]