FROM golang
WORKDIR /go/src/webserver
COPY webserver.go .
RUN GOOS=linux GOARCH=386 go build -a ./webserver.go

FROM scratch
WORKDIR /var/_data/
COPY --from=0 /go/src/webserver/webserver .
COPY . /var/_data/
EXPOSE 8080
CMD ["./webserver"]