FROM golang as build
WORKDIR /go/src/webserver
COPY ./ ./

RUN go get -u github.com/go-bindata/go-bindata/...
RUN go-bindata ./... 
RUN ls -l && mv webserver/webserver.go ./ && GOOS=linux GOARCH=386 go build -a ./*.go
RUN mv bindata web


FROM sbeliakou/kind:master-1.16.2
WORKDIR /var/_data/

COPY --from=build /go/src/webserver/web ./
COPY --from=build /go/src/webserver/main.html ./

COPY webserver/web.service /etc/systemd/system/
RUN systemctl enable web
EXPOSE 8081