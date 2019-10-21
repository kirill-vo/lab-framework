FROM golang
WORKDIR /go/src/webserver
COPY webserver.go .
RUN go get -u github.com/go-bindata/go-bindata/... && \
    go-bindata -ignore=\\webserver.go ./... && \
    GOOS=linux GOARCH=386 go build -a ./webserver.go

FROM scratch
WORKDIR /var/_data/
COPY --from=0 /go/src/webserver/* .
COPY . /var/_data/
EXPOSE 8080
CMD ["./webserver"]