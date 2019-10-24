FROM golang as build
WORKDIR /go/src/webserver
COPY ./ ./

RUN sed -i "s/###/$(ls -ld tasks/*/ | wc -l)/" webserver/webserver.go
RUN go get -u github.com/go-bindata/go-bindata/...
RUN go-bindata ./... 
RUN ls -l && mv webserver/webserver.go ./ && GOOS=linux GOARCH=386 go build -a ./*.go
RUN mv bindata web


FROM sbeliakou/kind:master-1.16.2
WORKDIR /var/_data/

RUN kubeadm reset cluster -f && rm -rf ~/.kube/*
RUN cat /kind/manifests/default-cni.yaml | sed 's/.. .PodSubnet ../10.244.0.0\/16/' > /opt/pod-network.yaml
RUN wget -O- https://raw.githubusercontent.com/kubernetes/kops/master/addons/metrics-server/v1.8.x.yaml | sed 's@apiVersion: extensions/v1beta1@apiVersion: apps/v1@' > /opt/metrics-server.yaml
RUN echo '#!/bin/bash\n\nexport HOME=/root\nexport KUBECONFIG=~/.kube/config\nexport TERM=xterm\nbash -i -l\n' > /usr/bin/bash-gotty

COPY --from=build /go/src/webserver/web ./
COPY --from=build /go/src/webserver/main.html ./

COPY webserver/web.service /etc/systemd/system/
RUN systemctl enable web

# RUN apt-get install -y sudo

ENV KUBECONFIG /root/.kube/config

RUN echo 

EXPOSE 8081