namespace = app
pull_policy = Always
project = client-go
name = http-api-server-check-deploy
prefix = $(namespace)-
uniq:=$(shell uuidgen)
stage = development
tag = $(shell whoami)-dev-$(uniq)

.PHONY: k8s_deploy
k8s_deploy: create_namespace app_build docker_image_build registry_push build_manifest kube_deploy

.PHONY: app_build
app_build: 
	GOOS=linux go build -o ./app .

.PHONY: docker_image_build
docker_image_build:
	docker build --tag $(name) .
	docker tag $(name):latest edmondandy/$(name):$(tag)

.PHONY: build_manifest
build_manifest:
	cat ./manifest_template.yaml | sed s/__PULL_POLICY__/$(pull_policy)/g | sed s/__NAMESPACE__/$(prefix)$(stage)/g | sed s/__STAGE__/$(stage)/g | sed s/__IMAGE__/$(name)/g | sed s/__NAME__/$(name)/g | sed s/__HASH_TAG__/$(tag)/ > ./manifest.yaml

.PHONY: registry_push
registry_push:
	docker push edmondandy/$(name):$(tag)

.PHONY: kube_deploy
kube_deploy:
	kubectl apply -f ./manifest.yaml

.PHONY: get_exposed_ip
get_exposed_ip:
	@echo "Deployment available at:"
	@kubectl get service $(name) -o json --namespace $(prefix)$(stage)| jq ".status.loadBalancer.ingress[].ip" | sed s/\"//g

.PHONY: teardown
teardown:
	kubectl delete namespace $(prefix)$(stage)

.PHONY: create_namespace
create_namespace:
	-kubectl create namespace $(prefix)$(stage)
