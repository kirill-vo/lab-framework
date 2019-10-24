FROM golang as build
WORKDIR /go/src/webserver
COPY ./ ./

RUN sed -i "s/###/$(ls -ld tasks/*/ | wc -l)/" webserver/webserver.go
RUN go get -u github.com/go-bindata/go-bindata/...
RUN go-bindata ./... 
RUN ls -l && mv webserver/webserver.go ./ && GOOS=linux GOARCH=386 go build -a ./*.go
RUN mv bindata web


FROM sbeliakou/kind:master-1.16.2
# FROM sbeliakou/kind:node-1.16.2
WORKDIR /var/_data/

COPY --from=build /go/src/webserver/web ./
COPY --from=build /go/src/webserver/main.html ./

COPY webserver/web.service /etc/systemd/system/
# RUN rm -rf /etc/kubernetes/*
RUN systemctl enable web

RUN echo '#!/bin/bash\n\nexport HOME=/root\nexport KUBECONFIG=~/.kube/config\nexport TERM=xterm\nbash -i -l\n' > /usr/bin/bash-gotty

# RUN apt-get install -y sudo

ENV KUBECONFIG /root/.kube/config

EXPOSE 8081