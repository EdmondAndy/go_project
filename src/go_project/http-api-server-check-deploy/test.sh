#!/bin/bash

curl http://localhost:8080/deployments

# result: [{"name":"nginx-deployment","namespace":"default"},{"name":"coredns","namespace":"kube-system"},{"name":"local-path-provisioner","namespace":"local-path-storage"}]