FROM golang as build
WORKDIR /go/src/webserver
COPY ./ ./

RUN ls -l && \
    find ./tasks -type d | sed '1d' | sort | xargs -IF bash -c "echo '-'; cd F; find ./ -type f | sed 's@^@F/@;s@\./@@g'" |  \
    awk 'BEGIN{print "courses:"}{if(/^-$/)print "-"; if(/verify/)print "  verify: ",$0; if(/task.md/)print "  task:   ",$0; if(/foreground/)print "  foreground: ", $0; if(/courseData/)print "  courseData: ", $0}' > course.yaml

RUN go get -u github.com/go-bindata/go-bindata/...
RUN go get github.com/smallfish/simpleyaml
RUN go-bindata ./... 
RUN ls -l && mv webserver/webserver.go ./ && GOOS=linux GOARCH=386 go build -a ./*.go
RUN mv bindata web


FROM sbeliakou/kind:master-1.16.2
WORKDIR /var/_data/

COPY --from=build /go/src/webserver/course.yaml ./
COPY --from=build /go/src/webserver/web ./

COPY webserver/web.service /etc/systemd/system/
RUN systemctl enable web
EXPOSE 8081