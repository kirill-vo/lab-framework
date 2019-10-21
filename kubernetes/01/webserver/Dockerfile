FROM golang
WORKDIR /go/src/webserver
COPY ./* ./
RUN go get -u github.com/go-bindata/go-bindata/... && \
    go-bindata ./... && \
    GOOS=linux GOARCH=386 go build -a ./*.go && \
    mv bindata web

FROM sbeliakou/kind:master-1.16.2
WORKDIR /var/_data/
COPY --from=0 /go/src/webserver/web ./
COPY web.service /etc/systemd/system/
RUN systemctl enable web
EXPOSE 8081

# FROM scratch
# WORKDIR /var/_data/
# COPY --from=0 /go/src/webserver/webserver ./
# EXPOSE 8080
# CMD ["./webserver"]