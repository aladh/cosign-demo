FROM golang:1.20

COPY . /build/

WORKDIR /build/

RUN go build -o /server .

ENTRYPOINT [ "/server" ]
