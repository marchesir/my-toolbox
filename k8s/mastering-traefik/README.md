colima start --kubernetes
1) whoami app has service created which ends up on clusterIP on k3s so not accessible.
2) whoami-ingressroute.yaml --> tradional ingress mapped to whoami.localhost.
3) whoami-httproute.yaml --> GatewayAPI but must be in same namespace as traefik and mapped to whoami-gatewayapi.localhost.
4) both ingress and httproute may not work at the same time as they both try to bind to same entrypoint (web).  httproute is prioritised if both are present.
5) create middle to intercept traffic to provide auth for dashboard.
6) create secrets from users
kubectl create secret generic dashboard-users --from-file=users -n traefik



