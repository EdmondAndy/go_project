#!/bin/bash

curl http://localhost:8080/replicas

# result: [{"name":"coredns","namespace":"kube-system","replicas":2,"ready":2,"available":2,"updated":2},{"name":"local-path-provisioner","namespace":"local-path-storage","replicas":1,"ready":1,"available":1,"updated":1},{"name":"nginx-deployment","namespace":"default","replicas":2,"ready":2,"available":2,"updated":2}]