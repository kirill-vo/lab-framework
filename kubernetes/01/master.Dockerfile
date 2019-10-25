FROM golang as build
WORKDIR /go/src/webserver
COPY ./ ./

RUN sed -i "s/\(var count_steps int = \).*/\1$(ls -ld tasks/*/ | wc -l)/" webserver/webserver.go
RUN go get -u github.com/go-bindata/go-bindata/...
RUN go-bindata ./... 
RUN ls -l && mv webserver/webserver.go ./ && GOOS=linux GOARCH=386 go build -a ./*.go
RUN mv bindata web


FROM sbeliakou/kind:master-1.15.3
WORKDIR /var/_data/

RUN kubeadm reset cluster -f && rm -rf ~/.kube/*
RUN echo '#!/bin/bash\n\nexport HOME=/root\nexport KUBECONFIG=~/.kube/config\nexport TERM=xterm\nbash -i -l\n' > /usr/bin/bash-gotty

COPY --from=build /go/src/webserver/web ./
COPY --from=build /go/src/webserver/main.html ./

COPY webserver/web.service /etc/systemd/system/
RUN systemctl enable web

ENV KUBECONFIG /root/.kube/config

RUN apt-get install -y sudo

EXPOSE 8081
EXPOSE 9090
EXPOSE 8001
EXPOSE 80
EXPOSE 443
EXPOSE 6443

WORKDIR /