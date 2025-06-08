#!/bin/bash

curl -X POST http://localhost:8080/scale   -H "Content-Type: application/json"   -d '{"namespace": "default", "deployment": "nginx-deployment", "replicas": 2}'

# result: Deployment nginx-deployment scaled to 2 replicas

curl -X POST http://localhost:8080/scale \
    -H "Content-Type: application/json" \
    -d '{"namespace": "default", "deployment": "my-app", "replicas": 5}'

# result: Deployment not found