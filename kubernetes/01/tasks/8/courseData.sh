#!/bin/bash

kubectl run green --restart=Never --image=adalimayeu/webapp-color --env="COLOR=green" --port=80
kubectl expose pod green --name=green-svc --port=80

cat << EOF | kubectl apply -f-
apiVersion: networking.k8s.io/v1beta1
kind: Ingress
metadata:
  name: green-ingress
  annotations:
    nginx.ingress.kubernetes.io/rewrite-target: /
spec:
  rules:
  - http:
      paths:
      - path: /green
        backend:
          serviceName: green-svc
          servicePort: 80
EOF


mkdir -p /opt/manifests/
cat << EOF > /opt/manifests/ingress-svc.yaml
kind: Service
apiVersion: v1
metadata:
  name: ingress-nginx
  namespace: ingress-nginx
  labels:
    app.kubernetes.io/name: ingress-nginx
    app.kubernetes.io/part-of: ingress-nginx
spec:
  clusterIP: 10.96.250.25
  externalIPs:
  - 172.31.0.2
  externalTrafficPolicy: Local
  healthCheckNodePort: 32511
  ports:
  - name: http
    nodePort: 31684
    port: 80
    protocol: TCP
    targetPort: http
  selector:
    app.kubernetes.io/name: ingress-nginx
    app.kubernetes.io/part-of: ingress-nginx
  sessionAffinity: None
  type: LoadBalancer
EOF



