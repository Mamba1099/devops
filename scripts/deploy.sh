#!/bin/bash

set -e

echo "🚀 Deploying to Kubernetes..."

# Deploy ConfigMap
kubectl apply -f k8s/base/configmap.yaml

# Deploy with Kustomize
kustomize build k8s/overlays/${1:-dev} | kubectl apply -f -

echo "✅ Deployment complete!"