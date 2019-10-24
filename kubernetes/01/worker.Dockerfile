FROM sbeliakou/kind:node-1.16.2
RUN kubeadm reset cluster -f
RUN rm -rf /etc/kubernetes/*