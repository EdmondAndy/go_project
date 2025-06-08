#!/bin/bash

curl "http://localhost:8080/replicas?namespace=default&name=nginx-deployment"

# result: {"name":"nginx-deployment","namespace":"default","desired_replicas":3,"available_replicas":3,"ready_replicas":3}

curl "http://localhost:8080/replicas?namespace=default&name=my-deployment"

# result: Error getting deployment: deployments.apps "my-deployment" not found