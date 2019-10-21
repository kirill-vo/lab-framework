FROM golang
WORKDIR /go/src/webserver
COPY ./* ./
RUN go get -u github.com/go-bindata/go-bindata/... && \
    go-bindata ./... && \
    GOOS=linux GOARCH=386 go build -a ./*.go && ls -l


FROM scratch
WORKDIR /var/_data/
COPY --from=0 /go/src/webserver/bindata .
EXPOSE 8080
CMD ["./bindata"]