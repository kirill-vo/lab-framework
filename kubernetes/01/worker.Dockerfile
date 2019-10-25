FROM sbeliakou/kind:node-1.15.3
RUN kubeadm reset cluster -f
RUN rm -rf /etc/kubernetes/*
WORKDIR /