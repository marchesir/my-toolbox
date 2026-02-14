#!/bin/bash

kubectl create ns traefik
helm repo add traefik https://traefik.github.io/charts
helm repo update
helm install --namespace=traefik  traefik traefik/traefik -f traefik-values.yaml --wait

# Creates traefik ns and installs traefik via traefik-values.yaml 
# -- installs dashboard at dashboard.localhost 
# -- installs GatewayAPI CRDs